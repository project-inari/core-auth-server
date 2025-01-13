package handler

import (
	"context"

	"github.com/project-inari/core-auth-server/protobuf/authPb"
)

// GRPCHandler represents the gRPC handler
type GRPCHandler struct {
	d Dependencies
	authPb.UnimplementedAuthGrpcServiceServer
}

func newGRPCHandler(d Dependencies) *GRPCHandler {
	return &GRPCHandler{d: d}
}

// VerifyToken verifies the token with adaptor-firebase-auth
func (h *GRPCHandler) VerifyToken(ctx context.Context, req *authPb.VerifyTokenReq) (*authPb.VerifyTokenRes, error) {
	res, err := h.d.Service.VerifyToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
