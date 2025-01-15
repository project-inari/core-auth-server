// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/project-inari/core-auth-server/dto"
	"github.com/project-inari/core-auth-server/protobuf/authPb"
	"github.com/project-inari/core-auth-server/repository"
)

// Port represents the service layer functions
type Port interface {
	SignUp(ctx context.Context, req dto.SignUpReq, h dto.SignUpReqHeader) (*dto.SignUpRes, error)
	VerifyToken(ctx context.Context, req *authPb.VerifyTokenReq) (*authPb.VerifyTokenRes, error)
	DeleteFirebaseUser(ctx context.Context, req *dto.DeleteFirebaseUserReq) (*dto.DeleteFirebaseUserRes, error)
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
