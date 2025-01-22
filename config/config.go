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
	APIAdaptorFirebaseAuthConfig APIAdaptorFirebaseAuthConfig
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

// APIAdaptorFirebaseAuthConfig represents the configuration of the adaptor-firebase-auth API
type APIAdaptorFirebaseAuthConfig struct {
	BaseURL                  string        `env:"API_ADAPTOR_FIREBASE_AUTH_BASE_URL,notEmpty"`
	SignupPath               string        `env:"API_ADAPTOR_FIREBASE_AUTH_SIGNUP_PATH,notEmpty"`
	VerifyTokenPath          string        `env:"API_ADAPTOR_FIREBASE_AUTH_VERIFY_TOKEN_PATH,notEmpty"`
	DeleteUserPath           string        `env:"API_ADAPTOR_FIREBASE_AUTH_DELETE_USER_PATH,notEmpty"`
	MaxConns                 int           `env:"API_ADAPTOR_FIREBASE_AUTH_MAX_CONNS,notEmpty"`
	MaxRetry                 int           `env:"API_ADAPTOR_FIREBASE_AUTH_MAX_RETRY"`
	Timeout                  time.Duration `env:"API_ADAPTOR_FIREBASE_AUTH_TIMEOUT,notEmpty"`
	InsecureSkipVerify       bool          `env:"API_ADAPTOR_FIREBASE_AUTH_INSECURE_SKIP_VERIFY,notEmpty"`
	MaxTransactionsPerSecond int           `env:"API_ADAPTOR_FIREBASE_AUTH_MAX_TRANSACTIONS_PER_SECOND"`
	DisableLogTrace          bool          `env:"API_ADAPTOR_FIREBASE_AUTH_DISABLE_LOG_TRACE"`
}
