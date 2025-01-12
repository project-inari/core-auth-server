package service

import (
	"context"

	"github.com/project-inari/core-auth-server/dto"
)

const (
	defaultAcceptLocale = "EN"
)

// SignUp calls the SignUp endpoint of the adaptor-firebase-auth service
func (s *service) SignUp(ctx context.Context, req dto.SignUpReq, h dto.SignUpReqHeader) (*dto.SignUpRes, error) {
	adaptorReq := dto.AdaptorFirebaseAuthSignUpReq(req)

	var acceptLocale string
	if h.AcceptLocale == "" {
		acceptLocale = defaultAcceptLocale
	} else {
		acceptLocale = h.AcceptLocale
	}

	adaptorHeader := dto.AdaptorFirebaseAuthSignUpReqHeader{
		AcceptLocale: acceptLocale,
	}

	res, err := s.adaptorFirebaseAuthRepository.CallSignUp(ctx, adaptorReq, adaptorHeader)
	if err != nil {
		return nil, err
	}

	return &dto.SignUpRes{
		Token: res.Response.Token,
	}, nil
}
