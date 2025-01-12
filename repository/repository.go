// Package repository provides the repository interfaces for the domain
package repository

import (
	"context"

	"github.com/project-inari/core-auth-server/dto"
	"github.com/project-inari/core-auth-server/pkg/httpclient"
)

// AdaptorFirebaseAuthRepository represents the repository API for the adaptor-firebase-auth service
type AdaptorFirebaseAuthRepository interface {
	CallSignUp(ctx context.Context, req dto.AdaptorFirebaseAuthSignUpReq, h dto.AdaptorFirebaseAuthSignUpReqHeader) (*httpclient.Response[dto.AdaptorFirebaseAuthSignUpRes], error)
}
