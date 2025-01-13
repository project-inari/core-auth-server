package di

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/project-inari/core-auth-server/config"
)

func setupGRPCServer(c *config.Config) (*grpc.Server, net.Listener) {
	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.AppConfig.GRPCPort))
	if err != nil {
		log.Panicf("error - [grpcconn.NewGrpcClient]: unable to listen on %v port: %v", c.AppConfig.GRPCPort, err)
	}

	return grpcServer, lis
}
