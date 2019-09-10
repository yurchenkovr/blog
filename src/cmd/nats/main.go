package main

import (
	"blog/src/infrastructure/config"
	gC "blog/src/repository/grpc/client"
	sub "blog/src/repository/nats/nats-sub"
	"blog/src/repository/postgres"
	"flag"
	"log"
)

// docker:  ./src/cmd/nats/grpcConfig.yaml
// local:  ./src/cmd/cmdmanager/grpcConfig.yaml
func main() {
	grpcPath := flag.String("p", "./src/cmd/nats/grpcConfig.yaml", "Path to gRPC config file")
	flag.Parse()

	cfg, err := config.Load(*grpcPath)
	if err != nil {
		log.Printf("Error while Loading config file\nReason: %v\n", err)
	}

	config := gC.Configs(cfg.Grpc.Host, cfg.Grpc.Port)

	dbHandler := postgres.New(config.APIms)

	sub.StartServer(dbHandler, config.NATSms)
}
