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
	"github.com/pawatOrbit/ai-mock-data-service/go/core/validation"
	"github.com/pawatOrbit/ai-mock-data-service/go/internal/model"
	"github.com/pawatOrbit/ai-mock-data-service/go/internal/repository"
	"github.com/pawatOrbit/ai-mock-data-service/go/utils"
)

const (
	ROLE_LM_STUDIO = "user"
)

type GenerateMockDataService interface {
	GenerateMockDataWithOneTable(ctx context.Context, in *model.GenerateMockDataWithOneTableRequest) (*model.GenerateMockDataWithOneTableResponse, error)
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
