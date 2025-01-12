// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/project-inari/core-auth-server/dto"
	"github.com/project-inari/core-auth-server/repository"
)

// Port represents the service layer functions
type Port interface {
	DoWiremock(ctx context.Context) (*dto.WiremockGetTestResponse, error)
}

type service struct {
	wiremockAPIRepository repository.WiremockAPIRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	WiremockAPIRepository repository.WiremockAPIRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		wiremockAPIRepository: d.WiremockAPIRepository,
	}
}
