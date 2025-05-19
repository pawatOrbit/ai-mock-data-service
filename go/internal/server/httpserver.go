package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/pawatOrbit/ai-mock-data-service/go/config"
	middleware_httpserver "github.com/pawatOrbit/ai-mock-data-service/go/core/transport/httpserver/middlewares"
	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewHttpServer() (*http.Server, error) {
	cfg := config.GetConfig()
	slog.InfoContext(context.Background(), "Initializing HTTP server", "port", cfg.RestServer.Port)
	var middlewares []middleware_httpserver.TransportMiddleware
	middlewares = append(middlewares, cors.New(cors.Options{
		AllowedOrigins: cfg.CORS.AllowedOrigins,
		AllowedMethods: cfg.CORS.AllowedMethods,
		AllowedHeaders: cfg.CORS.AllowedHeaders,
		ExposedHeaders: cfg.CORS.ExposedHeaders,
		MaxAge:         cfg.CORS.MaxAge,
	}).Handler)

	middlewareStack := middleware_httpserver.CreateStack(middlewares...)
	handler := registerRoute()
	wrappedMiddleware := middlewareStack(handler)
	wrappedOtel := otelhttp.NewHandler(
		wrappedMiddleware,
		"",
		otelhttp.WithSpanNameFormatter(
			func(operation string, r *http.Request) string {
				return fmt.Sprintf("%s %s %s", operation, r.Method, r.URL.Path)
			},
		))

	return &http.Server{
		Addr:    ":" + cfg.RestServer.Port,
		Handler: wrappedOtel,
	}, nil
}
