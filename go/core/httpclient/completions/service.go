package completions

import (
	"context"
	"net/http"
	"time"

	core_config "github.com/pawatOrbit/ai-mock-data-service/go/core/config"
	"github.com/pawatOrbit/ai-mock-data-service/go/core/httpclient/common"
)

type CompletionsServiceClient interface {
	GetCompletionsService(ctx context.Context, req CompletionRequest) (CompletionResponse, error)
}

type completionsServiceClient struct {
	cfg        *core_config.LMStudioConfig
	httpClient *http.Client
}

func NewCompletionsServiceClient(cfg *core_config.LMStudioConfig) CompletionsServiceClient {
	httpClient := http.Client{
		Timeout: 10 * time.Minute,
	}
	return &completionsServiceClient{
		cfg:        cfg,
		httpClient: &httpClient,
	}
}

func (s *completionsServiceClient) GetCompletionsService(ctx context.Context, req CompletionRequest) (CompletionResponse, error) {
	path := GET_COMPLETIONS_URL
	return common.Do[CompletionRequest, CompletionResponse, *CompletionError](ctx, s.cfg, s.httpClient, path, req)
}
