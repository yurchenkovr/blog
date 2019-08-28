package main

import (
	"blog/src/infrastructure/config"
	"blog/src/infrastructure/transport"
	gC "blog/src/usecases/grpc/client"
	"flag"
	"log"
)

func main() {
	grpcPath := flag.String("p", "./src/cmd/cmdmanager/grpcConfig.yaml", "Path to gRPC config file")
	flag.Parse()

	cfg, err := config.Load(*grpcPath)
	if err != nil {
		log.Printf("Error while Loading config file\nReason: %v\n", err)
	}

	config := gC.Configs(cfg.Grpc.Port)

	e := transport.New(config.APIms)
	e.Logger.Fatal(e.Start(config.APIms.Server.Port))
}
