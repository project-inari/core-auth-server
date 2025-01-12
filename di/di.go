// Package di provides dependency injection for the server
package di

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"

	"github.com/project-inari/core-auth-server/config"
	"github.com/project-inari/core-auth-server/handler"
	"github.com/project-inari/core-auth-server/pkg/httpclient"
	"github.com/project-inari/core-auth-server/repository"
	"github.com/project-inari/core-auth-server/service"
)

// New injects the dependencies for the server
func New(c *config.Config) {
	ctx := context.Background()

	// Sentry initialization
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:                c.SentryConfig.SentryDSN,
		Debug:              true,
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
	}); err != nil {
		slog.Error("error - [main.New] sentry initialization failed", slog.Any("error", err))
	}

	// Echo server initialization
	e := echo.New()
	setupServer(ctx, e, c)

	// HTTP Client initialization
	httpClientAdaptorFirebaseAuth := httpclient.NewHTTPClient(httpclient.Options{
		MaxConns:                 c.AdaptorFirebaseAuthAPIConfig.MaxConns,
		MaxRetry:                 c.AdaptorFirebaseAuthAPIConfig.MaxRetry,
		Timeout:                  c.AdaptorFirebaseAuthAPIConfig.Timeout,
		InsecureSkipVerify:       c.AdaptorFirebaseAuthAPIConfig.InsecureSkipVerify,
		MaxTransactionsPerSecond: c.AdaptorFirebaseAuthAPIConfig.MaxTransactionsPerSecond,
	})

	// Repository initialization
	adaptorFirebaseAuthRepo := repository.NewAdaptorFirebaseAuthRepository(repository.AdaptorFirebaseAuthRepositoryConfig{
		BaseURL: c.AdaptorFirebaseAuthAPIConfig.BaseURL,
		Path:    c.AdaptorFirebaseAuthAPIConfig.Path,
	}, repository.AdaptorFirebaseAuthRepositoryDependencies{
		Client: httpClientAdaptorFirebaseAuth,
	})

	// Service initialization
	service := service.New(service.Dependencies{
		AdaptorFirebaseAuthRepository: adaptorFirebaseAuthRepo,
	})

	// Handler initialization
	handler.New(e, handler.Dependencies{
		Service: service,
	})

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.New] unable to start server: %v", err)
	}
}
