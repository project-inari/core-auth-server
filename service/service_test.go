package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/project-inari/core-auth-server/dto"
	"github.com/project-inari/core-auth-server/pkg/httpclient"
	"github.com/project-inari/core-auth-server/protobuf/authPb"
)

type mockAdaptorFirebaseAuthRepository struct {
	signupRes      *httpclient.Response[dto.AdaptorFirebaseAuthSignUpRes]
	verifyTokenRes *httpclient.Response[dto.AdaptorFirebaseAuthVerifyTokenRes]
	deleteUserRes  *httpclient.Response[dto.AdaptorFirebaseAuthDeleteUserRes]
	err            error
}

func (m *mockAdaptorFirebaseAuthRepository) CallSignUp(_ context.Context, _ dto.AdaptorFirebaseAuthSignUpReq, _ dto.AdaptorFirebaseAuthSignUpReqHeader) (*httpclient.Response[dto.AdaptorFirebaseAuthSignUpRes], error) {
	return m.signupRes, m.err
}

func (m *mockAdaptorFirebaseAuthRepository) CallVerifyToken(_ context.Context, _ dto.AdaptorFirebaseAuthVerifyTokenReq, _ dto.AdaptorFirebaseAuthVerifyTokenReqHeader) (*httpclient.Response[dto.AdaptorFirebaseAuthVerifyTokenRes], error) {
	return m.verifyTokenRes, m.err
}

func (m *mockAdaptorFirebaseAuthRepository) CallDeleteUser(_ context.Context, _ dto.AdaptorFirebaseAuthDeleteUserReq, _ dto.AdaptorFirebaseAuthDeleteUserReqHeader) (*httpclient.Response[dto.AdaptorFirebaseAuthDeleteUserRes], error) {
	return m.deleteUserRes, m.err
}

const (
	mockToken    = "Bearer mockToken"
	mockUsername = "mockUsername"
	mockUID      = "mockUID"
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
				Username: mockUsername,
				UID:      mockUID,
				Token:    mockToken,
			},
		}

		adaptorFirebaseAuthRepository := &mockAdaptorFirebaseAuthRepository{
			signupRes: adaptorFirebaseAuthRes,
			err:       nil,
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
		assert.Equal(t, mockUsername, res.Username)
		assert.Equal(t, mockUID, res.UID)
		assert.Equal(t, mockToken, res.Token)
	})

	t.Run("error - when adaptor return status code not 200", func(t *testing.T) {
		adaptorFirebaseAuthRes := &httpclient.Response[dto.AdaptorFirebaseAuthSignUpRes]{
			HTTPStatusCode: 500,
			Response:       dto.AdaptorFirebaseAuthSignUpRes{},
		}

		adaptorFirebaseAuthRepository := &mockAdaptorFirebaseAuthRepository{
			signupRes: adaptorFirebaseAuthRes,
			err:       nil,
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

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("error - when httpclient error", func(t *testing.T) {
		adaptorFirebaseAuthRepository := &mockAdaptorFirebaseAuthRepository{
			signupRes: nil,
			err:       errors.New("error"),
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

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestVerifyToken(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		adaptorFirebaseAuthRes := &httpclient.Response[dto.AdaptorFirebaseAuthVerifyTokenRes]{
			HTTPStatusCode: 200,
			Response: dto.AdaptorFirebaseAuthVerifyTokenRes{
				Username: mockUsername,
				UID:      mockUID,
				Success:  true,
			},
		}

		adaptorFirebaseAuthRepository := &mockAdaptorFirebaseAuthRepository{
			verifyTokenRes: adaptorFirebaseAuthRes,
			err:            nil,
		}

		svc := &service{
			adaptorFirebaseAuthRepository: adaptorFirebaseAuthRepository,
		}

		req := &authPb.VerifyTokenReq{
			Token: mockToken,
		}

		res, err := svc.VerifyToken(ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, mockUsername, res.Username)
		assert.Equal(t, mockUID, res.Uid)
		assert.True(t, res.Success)
	})

	t.Run("error - when adaptor return status code not 200", func(t *testing.T) {
		adaptorFirebaseAuthRes := &httpclient.Response[dto.AdaptorFirebaseAuthVerifyTokenRes]{
			HTTPStatusCode: 500,
			Response:       dto.AdaptorFirebaseAuthVerifyTokenRes{},
		}

		adaptorFirebaseAuthRepository := &mockAdaptorFirebaseAuthRepository{
			verifyTokenRes: adaptorFirebaseAuthRes,
			err:            nil,
		}

		svc := &service{
			adaptorFirebaseAuthRepository: adaptorFirebaseAuthRepository,
		}

		req := &authPb.VerifyTokenReq{
			Token: mockToken,
		}

		_, err := svc.VerifyToken(ctx, req)

		assert.NotNil(t, err)
	})

	t.Run("error - when httpclient error", func(t *testing.T) {
		adaptorFirebaseAuthRepository := &mockAdaptorFirebaseAuthRepository{
			verifyTokenRes: nil,
			err:            errors.New("error"),
		}

		svc := &service{
			adaptorFirebaseAuthRepository: adaptorFirebaseAuthRepository,
		}

		req := &authPb.VerifyTokenReq{
			Token: mockToken,
		}

		_, err := svc.VerifyToken(ctx, req)

		assert.NotNil(t, err)
	})
}

func TestDeleteFirebaseUser(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		adaptorFirebaseAuthRes := &httpclient.Response[dto.AdaptorFirebaseAuthDeleteUserRes]{
			HTTPStatusCode: 200,
			Response: dto.AdaptorFirebaseAuthDeleteUserRes{
				Success: true,
			},
		}

		adaptorFirebaseAuthRepo := &mockAdaptorFirebaseAuthRepository{
			deleteUserRes: adaptorFirebaseAuthRes,
			err:           nil,
		}

		svc := &service{
			adaptorFirebaseAuthRepository: adaptorFirebaseAuthRepo,
		}

		req := dto.DeleteFirebaseUserReq{
			UID: mockUID,
		}

		res, err := svc.DeleteFirebaseUser(ctx, req)

		assert.Nil(t, err)
		assert.True(t, res.Success)
	})

	t.Run("error - when adaptor return status code not 200", func(t *testing.T) {
		adaptorFirebaseAuthRes := &httpclient.Response[dto.AdaptorFirebaseAuthDeleteUserRes]{
			HTTPStatusCode: 500,
			Response:       dto.AdaptorFirebaseAuthDeleteUserRes{},
		}

		adaptorFirebaseAuthRepo := &mockAdaptorFirebaseAuthRepository{
			deleteUserRes: adaptorFirebaseAuthRes,
			err:           nil,
		}

		svc := &service{
			adaptorFirebaseAuthRepository: adaptorFirebaseAuthRepo,
		}

		req := dto.DeleteFirebaseUserReq{
			UID: mockUID,
		}

		_, err := svc.DeleteFirebaseUser(ctx, req)

		assert.NotNil(t, err)
	})

	t.Run("error - when httpclient error", func(t *testing.T) {
		adaptorFirebaseAuthRepo := &mockAdaptorFirebaseAuthRepository{
			deleteUserRes: nil,
			err:           errors.New("error"),
		}

		svc := &service{
			adaptorFirebaseAuthRepository: adaptorFirebaseAuthRepo,
		}

		req := dto.DeleteFirebaseUserReq{
			UID: mockUID,
		}

		_, err := svc.DeleteFirebaseUser(ctx, req)

		assert.NotNil(t, err)
	})
}
