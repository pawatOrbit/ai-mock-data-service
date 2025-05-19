package server

import (
	"context"
	"net/http"

	"github.com/pawatOrbit/ai-mock-data-service/go/core/transport/httpserver"
	middleware_httpserver "github.com/pawatOrbit/ai-mock-data-service/go/core/transport/httpserver/middlewares"
	"github.com/pawatOrbit/ai-mock-data-service/go/internal/model"
)

func registerRoute() http.Handler {
	mux := http.NewServeMux()
	r := httpserver.NewRouter(mux)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware_httpserver.NotFound(w, r)
	}))

	r.Post("/health-check",
		httpserver.NewTransport(
			&model.HealthReq{},
			httpserver.NewEndpoint(func(ctx context.Context, in *model.HealthReq) (*model.HealthResp, error) {
				return &model.HealthResp{
					Status:   1000,
					Response: "Hello, " + in.Name,
				}, nil
			})))

	return mux
}
