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
	CallVerifyToken(ctx context.Context, req dto.AdaptorFirebaseAuthVerifyTokenReq, h dto.AdaptorFirebaseAuthVerifyTokenReqHeader) (*httpclient.Response[dto.AdaptorFirebaseAuthVerifyTokenRes], error)
	CallDeleteUser(ctx context.Context, req dto.AdaptorFirebaseAuthDeleteUserReq, h dto.AdaptorFirebaseAuthDeleteUserReqHeader) (*httpclient.Response[dto.AdaptorFirebaseAuthDeleteUserRes], error)
}
