package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/project-inari/core-auth-server/dto"
)

// DeleteFirebaseUser deletes a user in Firebase Auth
func (s *service) DeleteFirebaseUser(ctx context.Context, req *dto.DeleteFirebaseUserReq) (*dto.DeleteFirebaseUserRes, error) {
	res, err := s.adaptorFirebaseAuthRepository.CallDeleteUser(ctx, dto.AdaptorFirebaseAuthDeleteUserReq{
		UID: req.UID,
	}, dto.AdaptorFirebaseAuthDeleteUserReqHeader{})
	if err != nil {
		return nil, err
	}
	if res.HTTPStatusCode != http.StatusOK {
		return nil, fmt.Errorf("error - [service.DeleteFirebaseUser] - adaptor-firebase-auth http status code returns %v", res.HTTPStatusCode)
	}

	return &dto.DeleteFirebaseUserRes{
		Success: res.Response.Success,
	}, nil
}
