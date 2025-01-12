package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/project-inari/core-auth-server/dto"
	"github.com/project-inari/core-auth-server/pkg/httpclient"
)

type adaptorFirebaseAuthRepository struct {
	baseURL string
	path    string
	client  *http.Client
}

// AdaptorFirebaseAuthRepositoryConfig represents the configuration for the adaptorFirebaseAuthRepository
type AdaptorFirebaseAuthRepositoryConfig struct {
	BaseURL string
	Path    string
}

// AdaptorFirebaseAuthRepositoryDependencies represents the dependencies for the adaptorFirebaseAuthRepository
type AdaptorFirebaseAuthRepositoryDependencies struct {
	Client *http.Client
}

// NewAdaptorFirebaseAuthRepository creates a new adaptorFirebaseAuthRepository
func NewAdaptorFirebaseAuthRepository(c AdaptorFirebaseAuthRepositoryConfig, d AdaptorFirebaseAuthRepositoryDependencies) AdaptorFirebaseAuthRepository {
	return &adaptorFirebaseAuthRepository{
		baseURL: c.BaseURL,
		path:    c.Path,
		client:  d.Client,
	}
}

// CallSignUp calls the Signup endpoint of the adaptor-firebase-auth service
func (r *adaptorFirebaseAuthRepository) CallSignUp(ctx context.Context, req dto.AdaptorFirebaseAuthSignUpReq, h dto.AdaptorFirebaseAuthSignUpReqHeader) (*httpclient.Response[dto.AdaptorFirebaseAuthSignUpRes], error) {
	url := fmt.Sprintf("%s%s", r.baseURL, r.path)
	return httpclient.Post[dto.AdaptorFirebaseAuthSignUpReq, dto.AdaptorFirebaseAuthSignUpRes](ctx, r.client, url, h.ToMap(), req)
}
