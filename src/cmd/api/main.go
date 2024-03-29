package main

import (
	"blog/src/infrastructure/config"
	"blog/src/infrastructure/transport"
	gC "blog/src/repository/grpc/client"
	"flag"
	"log"
)

// docker:  ./src/cmd/api/grpcConfig.yaml
// local:  ./src/cmd/cmdmanager/grpcConfig.yaml
func main() {
	grpcPath := flag.String("p", "./src/cmd/api/grpcConfig.yaml", "Path to gRPC config file")
	flag.Parse()

	cfg, err := config.Load(*grpcPath)
	if err != nil {
		log.Printf("Error while Loading config file\nReason: %v\n", err)
	}

	config := gC.Configs(cfg.Grpc.Host, cfg.Grpc.Port)

	e := transport.New(config.APIms, config.NATSms)
	e.Logger.Fatal(e.Start(config.APIms.Server.Port))
}
