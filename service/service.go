// Package service provides the business logic service layer for the server
package service

import (
	// "context"

	// "github.com/project-inari/core-auth-server/dto"
	"github.com/project-inari/core-auth-server/repository"
)

// Port represents the service layer functions
type Port interface {
}

type service struct {
	adaptorFirebaseAuthRepository repository.AdaptorFirebaseAuthRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	AdaptorFirebaseAuthRepository repository.AdaptorFirebaseAuthRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		adaptorFirebaseAuthRepository: d.AdaptorFirebaseAuthRepository,
	}
}
