package client

import (
	"blog/src/infrastructure/config"
	"blog/src/repository/grpc/routeconfig"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"log"
)

func StartClient(host, port string) *routeconfig.RouteConfigClient {
	conn, err := grpc.Dial(host+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v ", err)
	}

	c := routeconfig.NewRouteConfigClient(conn)

	return &c
}

func LoadConfigs(host, port string) *config.Configuration {
	c := *StartClient(host, port)

	r, err := c.GetServerConfig(context.Background(), &routeconfig.RequestName{})
	if err != nil {
		log.Fatalf("failed to getResponse: %v", err)
	}

	var config config.Configuration
	if err := json.Unmarshal(r.GetConfig(), &config); err != nil {
		log.Printf("error when unmarshaling config: %v", err)
		return nil
	}

	return &config
}

func Configs(host, port string) *config.Configuration {
	c := LoadConfigs(host, port)

	return c
}
