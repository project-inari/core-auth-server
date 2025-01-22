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
	"github.com/project-inari/core-auth-server/protobuf/authPb"
	"github.com/project-inari/core-auth-server/repository"
	"github.com/project-inari/core-auth-server/service"
)

// New injects the dependencies for the server
func New(c *config.Config) {
	ctx := context.Background()

	// Sentry initialization
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:                c.SentryConfig.SentryDSN,
		Debug:              false,
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
	}); err != nil {
		slog.Error("error - [main.New] sentry initialization failed", slog.Any("error", err))
	}

	// Echo server initialization
	e := echo.New()
	setupServer(ctx, e, c)

	// GRPC Server initialization
	grpcServer, grpcListener := setupGRPCServer(c)

	// HTTP Client initialization
	httpClientAdaptorFirebaseAuth := httpclient.NewHTTPClient(httpclient.Options{
		MaxConns:                 c.APIAdaptorFirebaseAuthConfig.MaxConns,
		MaxRetry:                 c.APIAdaptorFirebaseAuthConfig.MaxRetry,
		Timeout:                  c.APIAdaptorFirebaseAuthConfig.Timeout,
		InsecureSkipVerify:       c.APIAdaptorFirebaseAuthConfig.InsecureSkipVerify,
		MaxTransactionsPerSecond: c.APIAdaptorFirebaseAuthConfig.MaxTransactionsPerSecond,
		DisableLogTrace:          c.APIAdaptorFirebaseAuthConfig.DisableLogTrace,
	})

	// Repository initialization
	adaptorFirebaseAuthRepo := repository.NewAdaptorFirebaseAuthRepository(repository.AdaptorFirebaseAuthRepositoryConfig{
		BaseURL:         c.APIAdaptorFirebaseAuthConfig.BaseURL,
		SignupPath:      c.APIAdaptorFirebaseAuthConfig.SignupPath,
		VerifyTokenPath: c.APIAdaptorFirebaseAuthConfig.VerifyTokenPath,
		DeleteUserPath:  c.APIAdaptorFirebaseAuthConfig.DeleteUserPath,
	}, repository.AdaptorFirebaseAuthRepositoryDependencies{
		Client: httpClientAdaptorFirebaseAuth,
	})

	// Service initialization
	service := service.New(service.Dependencies{
		AdaptorFirebaseAuthRepository: adaptorFirebaseAuthRepo,
	})

	// Handler initialization
	grpcHandler := handler.New(e, handler.Dependencies{
		Service: service,
	})

	go func() {
		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)
		slog.Info("Starting gRPC server...", slog.Any("port", c.AppConfig.GRPCPort))
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Panicf("error - [handler.New]: unable to start gRPC server: %v", err)
		}
	}()

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.New] unable to start server: %v", err)
	}
}
