package service

import (
	"context"
	"strings"

	"github.com/project-inari/core-auth-server/dto"
	"github.com/project-inari/core-auth-server/protobuf/authPb"
)

const (
	bearerTokenPrefix = "Bearer "
)

func (s *service) VerifyToken(ctx context.Context, req *authPb.VerifyTokenReq) (*authPb.VerifyTokenRes, error) {
	res, err := s.adaptorFirebaseAuthRepository.CallVerifyToken(ctx, dto.AdaptorFirebaseAuthVerifyTokenReq{
		Token: strings.Split(req.Token, bearerTokenPrefix)[1],
	}, dto.AdaptorFirebaseAuthVerifyTokenReqHeader{})
	if err != nil {
		return nil, err
	}

	return &authPb.VerifyTokenRes{
		Username: res.Response.Username,
		Uid:      res.Response.UID,
		Success:  res.Response.Success,
	}, nil
}
