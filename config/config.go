// Package config provides configuration settings for the server
package config

import (
	"log"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var once sync.Once
var config *Config

// New loads the configuration from the .env file
func New() *Config {
	once.Do(func() {
		e := os.Getenv("APP_ENV_STAGE")
		if e == "" || e == "LOCAL" {
			if err := godotenv.Load(".env.generated"); err != nil {
				slog.Warn("[config.New] unable to load .env.generated file", slog.Any("error", err))
			}
		}

		cfg := &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("error - [config.New] unable to parse config: %v", err)
		}
		config = cfg
	})

	return config
}

// Config represents the configuration of the server
type Config struct {
	AppConfig                    AppConfig
	LogConfig                    LogConfig
	SentryConfig                 SentryConfig
	AdaptorFirebaseAuthAPIConfig AdaptorFirebaseAuthAPIConfig
}

// AppConfig represents the configuration of the application
type AppConfig struct {
	Name     string `env:"APP_NAME,notEmpty"`
	Port     string `env:"APP_PORT,notEmpty"`
	GRPCPort string `env:"APP_GRPC_PORT,notEmpty"`
	EnvStage string `env:"APP_ENV_STAGE,notEmpty"`
}

// LogConfig represents the configuration of the logger
type LogConfig struct {
	Level             string `env:"LOG_LEVEL,notEmpty"`
	MaskSensitiveData bool   `env:"LOG_MASK_SENSITIVE_DATA,notEmpty"`
}

// SentryConfig represents the configuration of Sentry.io
type SentryConfig struct {
	SentryDSN string `env:"SENTRY_DSN"`
}

// AdaptorFirebaseAuthAPIConfig represents the configuration of the adaptor-firebase-auth API
type AdaptorFirebaseAuthAPIConfig struct {
	BaseURL                  string        `env:"ADAPTOR_FIREBASE_AUTH_API_BASE_URL,notEmpty"`
	SignupPath               string        `env:"ADAPTOR_FIREBASE_AUTH_API_SIGNUP_PATH,notEmpty"`
	VerifyTokenPath          string        `env:"ADAPTOR_FIREBASE_AUTH_API_VERIFY_TOKEN_PATH,notEmpty"`
	MaxConns                 int           `env:"ADAPTOR_FIREBASE_AUTH_API_MAX_CONNS,notEmpty"`
	MaxRetry                 int           `env:"ADAPTOR_FIREBASE_AUTH_API_MAX_RETRY,notEmpty"`
	Timeout                  time.Duration `env:"ADAPTOR_FIREBASE_AUTH_API_TIMEOUT,notEmpty"`
	InsecureSkipVerify       bool          `env:"ADAPTOR_FIREBASE_AUTH_API_INSECURE_SKIP_VERIFY,notEmpty"`
	MaxTransactionsPerSecond int           `env:"ADAPTOR_FIREBASE_AUTH_API_MAX_TRANSACTIONS_PER_SECOND"`
	DisableLogTrace          bool          `env:"ADAPTOR_FIREBASE_AUTH_API_DISABLE_LOG_TRACE"`
}
