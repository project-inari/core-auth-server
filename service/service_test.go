package service

import (
	"context"
	"errors"
	"testing"

	"github.com/project-inari/core-auth-server/dto"
	"github.com/project-inari/core-auth-server/pkg/httpclient"
	"github.com/stretchr/testify/assert"
)

type mockAdaptorFirebaseAuthRepository struct {
	res *httpclient.Response[dto.AdaptorFirebaseAuthSignUpRes]
	err error
}

func (m *mockAdaptorFirebaseAuthRepository) CallSignUp(_ context.Context, _ dto.AdaptorFirebaseAuthSignUpReq, _ dto.AdaptorFirebaseAuthSignUpReqHeader) (*httpclient.Response[dto.AdaptorFirebaseAuthSignUpRes], error) {
	return m.res, m.err
}

const (
	mockToken    = "mockToken"
	mockUsername = "mockUsername"
	mockEmail    = "mock@email.com"
	mockPassword = "mockPassword"
	mockPhoneNo  = "+66000000000"

	enAcceptLocale = "EN"
	thAcceptLocale = "TH"
)

func TestSignUp(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		adaptorFirebaseAuthRes := &httpclient.Response[dto.AdaptorFirebaseAuthSignUpRes]{
			HTTPStatusCode: 200,
			Response: dto.AdaptorFirebaseAuthSignUpRes{
				Token: mockToken,
			},
		}

		adaptorFirebaseAuthRepository := &mockAdaptorFirebaseAuthRepository{
			res: adaptorFirebaseAuthRes,
			err: nil,
		}

		svc := &service{
			adaptorFirebaseAuthRepository: adaptorFirebaseAuthRepository,
		}

		req := dto.SignUpReq{
			Username: mockUsername,
			Email:    mockEmail,
			Password: mockPassword,
			PhoneNo:  mockPhoneNo,
		}
		header := dto.SignUpReqHeader{
			AcceptLocale: enAcceptLocale,
		}

		res, err := svc.SignUp(ctx, req, header)

		assert.Nil(t, err)
		assert.Equal(t, mockToken, res.Token)
	})

	t.Run("error - when adaptor return status code not 200", func(t *testing.T) {
		adaptorFirebaseAuthRes := &httpclient.Response[dto.AdaptorFirebaseAuthSignUpRes]{
			HTTPStatusCode: 500,
			Response:       dto.AdaptorFirebaseAuthSignUpRes{},
		}

		adaptorFirebaseAuthRepository := &mockAdaptorFirebaseAuthRepository{
			res: adaptorFirebaseAuthRes,
			err: nil,
		}

		svc := &service{
			adaptorFirebaseAuthRepository: adaptorFirebaseAuthRepository,
		}

		req := dto.SignUpReq{
			Username: mockUsername,
			Email:    mockEmail,
			Password: mockPassword,
			PhoneNo:  mockPhoneNo,
		}
		header := dto.SignUpReqHeader{
			AcceptLocale: enAcceptLocale,
		}

		_, err := svc.SignUp(ctx, req, header)

		assert.NotNil(t, err)
	})

	t.Run("error - when httpclient error", func(t *testing.T) {
		adaptorFirebaseAuthRepository := &mockAdaptorFirebaseAuthRepository{
			res: nil,
			err: errors.New("error"),
		}

		svc := &service{
			adaptorFirebaseAuthRepository: adaptorFirebaseAuthRepository,
		}

		req := dto.SignUpReq{
			Username: mockUsername,
			Email:    mockEmail,
			Password: mockPassword,
			PhoneNo:  mockPhoneNo,
		}
		header := dto.SignUpReqHeader{
			AcceptLocale: enAcceptLocale,
		}

		_, err := svc.SignUp(ctx, req, header)

		assert.NotNil(t, err)
	})
}
