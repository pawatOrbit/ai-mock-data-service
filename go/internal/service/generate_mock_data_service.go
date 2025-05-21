package service

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pawatOrbit/ai-mock-data-service/go/config"
	"github.com/pawatOrbit/ai-mock-data-service/go/core/exception"
	"github.com/pawatOrbit/ai-mock-data-service/go/core/httpclient"
	"github.com/pawatOrbit/ai-mock-data-service/go/core/httpclient/completions"
	"github.com/pawatOrbit/ai-mock-data-service/go/core/logger"
	"github.com/pawatOrbit/ai-mock-data-service/go/core/validation"
	"github.com/pawatOrbit/ai-mock-data-service/go/internal/model"
	"github.com/pawatOrbit/ai-mock-data-service/go/internal/repository"
	"github.com/pawatOrbit/ai-mock-data-service/go/utils"
	"github.com/pawatOrbit/ai-mock-data-service/go/utils/conv"
)

const (
	ROLE_LM_STUDIO = "user"
)

type GenerateMockDataService interface {
	GenerateMockDataWithOneTable(ctx context.Context, in *model.GenerateMockDataWithOneTableRequest) (*model.GenerateMockDataWithOneTableResponse, error)
	GenerateMockDataWithFkTables(ctx context.Context, in *model.GenerateMockDataWithFkTableRequest) (*model.GenerateMockDataWithFkTableResponse, error)
}

type generateMockDataService struct {
	Repo           *repository.Repository
	Errors         *exception.MockDataServiceErrors
	Utils          *utils.Utils
	Config         *config.Config
	LmStudioClient *httpclient.LmStudioServiceClient
}

func NewGenerateMockDataService(repo *repository.Repository, errors *exception.MockDataServiceErrors, utils *utils.Utils, config *config.Config, lmStudioClient *httpclient.LmStudioServiceClient) GenerateMockDataService {
	return &generateMockDataService{
		Repo:           repo,
		Errors:         errors,
		Utils:          utils,
		Config:         config,
		LmStudioClient: lmStudioClient,
	}
}

func (g *generateMockDataService) GenerateMockDataWithOneTable(ctx context.Context, in *model.GenerateMockDataWithOneTableRequest) (*model.GenerateMockDataWithOneTableResponse, error) {
	errValidate := validation.ValidateStruct(in)
	if errValidate != nil {
		return nil, g.Errors.ErrInvalidRequest.WithDebugMessage(errValidate.Error())
	}

	dateTimeNow := time.Now()

	// Get table schema
	tableSchema, err := g.Repo.TableSchemasRepository.GetDatabaseSchemaByTableName(ctx, in.TableName)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, g.Errors.ErrUnableToProceed.WithDebugMessage(err.Error())
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, g.Errors.ErrNotFound.WithDebugMessage("table not found")
	}

	prompt := g.Utils.GeneratePromptUtils.GeneratePromptWithoutKey(tableSchema.TableName, tableSchema.TableScript.String, in.NumSample)

	reqGetCompletionsService := completions.CompletionRequest{
		Model: g.Config.LMStudio.Model,
		Messages: []completions.MessageRequest{
			{
				Role:    ROLE_LM_STUDIO,
				Content: prompt,
			},
		},
		Temperature: g.Config.LMStudio.Temperature,
		MaxTokens:   g.Config.LMStudio.MaxTokens,
	}
	respLmStudio, err := g.LmStudioClient.GetCompletionsService.GetCompletionsService(ctx, reqGetCompletionsService)
	if err != nil {
		return nil, g.Errors.ErrUnableToProceed.WithDebugMessage(err.Error())
	}

	resp := &model.GenerateMockDataWithOneTableResponse{
		Status: 200,
		Data: model.GenerateMockDataWithOneTableResponseData{
			Query:            respLmStudio.Choices[0].Message.Content,
			PromptTokens:     respLmStudio.Usage.PromptTokens,
			CompletionTokens: respLmStudio.Usage.CompletionTokens,
			TotalTokens:      respLmStudio.Usage.TotalTokens,
			TimeTaken:        time.Since(dateTimeNow).Seconds(),
		},
	}

	return resp, nil
}

func (g *generateMockDataService) GenerateMockDataWithFkTables(ctx context.Context, in *model.GenerateMockDataWithFkTableRequest) (*model.GenerateMockDataWithFkTableResponse, error) {
	errValidate := validation.ValidateStruct(in)
	if errValidate != nil {
		return nil, g.Errors.ErrInvalidRequest.WithDebugMessage(errValidate.Error())
	}

	slog := logger.NewPathfinder("GenerateMockDataWithFkTables")

	dateTimeNow := time.Now()

	PromptTokens := 0
	CompletionTokens := 0
	TotalTokens := 0

	// Get table schema
	tableSchema, err := g.Repo.TableSchemasRepository.GetDatabaseSchemaByTableName(ctx, in.TableName)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, g.Errors.ErrUnableToProceed.WithDebugMessage(err.Error())
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, g.Errors.ErrNotFound.WithDebugMessage("table not found")
	}

	extractKeyPrompt := g.Utils.GeneratePromptUtils.GeneratePromptForFKExtraction(tableSchema.TableName, tableSchema.TableScript.String)
	reqGetCompletionsService := completions.CompletionRequest{
		Model: g.Config.LMStudio.Model,
		Messages: []completions.MessageRequest{
			{
				Role:    ROLE_LM_STUDIO,
				Content: extractKeyPrompt,
			},
		},
		Temperature: g.Config.LMStudio.Temperature,
		MaxTokens:   g.Config.LMStudio.MaxTokens,
	}

	respKeyExtraction, err := g.LmStudioClient.GetCompletionsService.GetCompletionsService(ctx, reqGetCompletionsService)
	if err != nil {
		return nil, g.Errors.ErrUnableToProceed.WithDebugMessage(err.Error())
	}

	PromptTokens += respKeyExtraction.Usage.PromptTokens
	CompletionTokens += respKeyExtraction.Usage.CompletionTokens
	TotalTokens += respKeyExtraction.Usage.TotalTokens

	_, fk := g.Utils.ExtractStringUtils.ExtractForeignKeyInfo(respKeyExtraction.Choices[0].Message.Content)

	insertResponse := []string{}
	fieldName := []string{}
	fieldValue := []string{}

	// index := 0
	for fkKey, v := range fk {
		fkTable := v[0]
		fkField := v[1]

		slog.DebugContext(ctx, "Foreign key", "fkKey", fkKey, "fkTable", fkTable, "fkField", fkField)

		tableSchemaFk, err := g.Repo.TableSchemasRepository.GetDatabaseSchemaByTableName(ctx, fkTable)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, g.Errors.ErrUnableToProceed.WithDebugMessage(err.Error())
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, g.Errors.ErrNotFound.WithDebugMessage("table not found")
		}

		extractKeyPrompt := g.Utils.GeneratePromptUtils.GeneratePromptWithoutKey(tableSchemaFk.TableName, tableSchemaFk.TableScript.String, 1)

		reqGetCompletionsService := completions.CompletionRequest{
			Model: g.Config.LMStudio.Model,
			Messages: []completions.MessageRequest{
				{
					Role:    ROLE_LM_STUDIO,
					Content: extractKeyPrompt,
				},
			},
			Temperature: g.Config.LMStudio.Temperature,
			MaxTokens:   g.Config.LMStudio.MaxTokens,
		}

		respFromFkTable, err := g.LmStudioClient.GetCompletionsService.GetCompletionsService(ctx, reqGetCompletionsService)
		if err != nil {
			return nil, g.Errors.ErrUnableToProceed.WithDebugMessage(err.Error())
		}

		PromptTokens += respFromFkTable.Usage.PromptTokens
		CompletionTokens += respFromFkTable.Usage.CompletionTokens
		TotalTokens += respFromFkTable.Usage.TotalTokens

		responseFromFkTable := conv.ReplaceNewlineWithSpace(respFromFkTable.Choices[0].Message.Content)

		insertResponse = append(insertResponse, responseFromFkTable)

		insertValueAndField, err := g.Utils.ExtractStringUtils.ExtractInsertValues(responseFromFkTable)
		if err != nil {
			return nil, g.Errors.ErrUnableToProceed.WithDebugMessage(err.Error())
		}
		if len(insertValueAndField) == 0 {
			return nil, g.Errors.ErrNotFound.WithDebugMessage("table not found")
		}
		fieldName = append(fieldName, fkKey)
		fieldValue = append(fieldValue, insertValueAndField[fkField])

	}

	prompt := g.Utils.GeneratePromptUtils.GeneratePromptForMockDataWithValues(tableSchema.TableName, tableSchema.TableScript.String, in.NumSample, fieldName, fieldValue)
	reqGetCompletionsMain := completions.CompletionRequest{
		Model: g.Config.LMStudio.Model,
		Messages: []completions.MessageRequest{
			{
				Role:    ROLE_LM_STUDIO,
				Content: prompt,
			},
		},
		Temperature: g.Config.LMStudio.Temperature,
		MaxTokens:   g.Config.LMStudio.MaxTokens,
	}

	respMain, err := g.LmStudioClient.GetCompletionsService.GetCompletionsService(ctx, reqGetCompletionsMain)
	if err != nil {
		return nil, g.Errors.ErrUnableToProceed.WithDebugMessage(err.Error())
	}

	PromptTokens += respMain.Usage.PromptTokens
	CompletionTokens += respMain.Usage.CompletionTokens
	TotalTokens += respMain.Usage.TotalTokens

	responseInsert := conv.ReplaceNewlineWithSpace(respMain.Choices[0].Message.Content) + " " + conv.JoinStringSlice(insertResponse, " ")

	slog.DebugContext(ctx, "Response insert", "responseInsert", responseInsert)

	return &model.GenerateMockDataWithFkTableResponse{
		Status: 200,
		Data: model.GenerateMockDataWithFkTableResponseData{
			Query:            responseInsert,
			PromptTokens:     PromptTokens,
			CompletionTokens: CompletionTokens,
			TotalTokens:      TotalTokens,
			TimeTaken:        time.Since(dateTimeNow).Seconds(),
		},
	}, nil

}
