package handler

import (
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
