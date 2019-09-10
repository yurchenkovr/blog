package main

import (
	"blog/src/infrastructure/config"
	gS "blog/src/repository/grpc/server"
	"flag"
	"log"
)

//  docker: ./src/cmd/cmdmanager/config.local.yaml
//  local: ./src/cmd/api/config.local.yaml
func main() {
	grpcPath := flag.String("p", "./src/cmd/cmdmanager/grpcConfig.yaml", "Path to gRPC config file")
	cfgPath := flag.String("cfgPath", "./src/cmd/cmdmanager/config.local.yaml", "Path to config file")
	flag.Parse()

	gcfg, err := config.Load(*grpcPath)
	cfg, err := config.Load(*cfgPath)

	server := gS.NewRouteConfigServer(*cfg)
	if err != nil {
		log.Printf("Error while Loading gRPC config file\nReason: %v\n", err)
	}

	gS.StartServer(gcfg.Grpc.Port, server)
}
