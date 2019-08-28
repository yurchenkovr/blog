package server

import (
	"blog/src/infrastructure/config"
	pb "blog/src/usecases/grpc/routeconfig"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net"
)

type RouteConfigServer struct {
	c *config.Configuration
}

func NewRouteConfigServer(c config.Configuration) *RouteConfigServer {
	return &RouteConfigServer{c: &c}
}

func (s *RouteConfigServer) GetServerConfig(ctx context.Context, in *pb.RequestName) (*pb.ServerConfig, error) {
	bytes, err := json.Marshal(&s.c)
	if err != nil {
		return nil, err
	}

	return &pb.ServerConfig{Config: bytes}, nil
}

func StartServer(port string, server *RouteConfigServer) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("gRPC Server was started.\n")

	grpcServer := grpc.NewServer()

	pb.RegisterRouteConfigServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v ", err)
	}
}
