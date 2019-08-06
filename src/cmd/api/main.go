package main

import (
	"blog/src/infrastructure/config"
	"blog/src/infrastructure/transport"
	"flag"
	"github.com/labstack/gommon/log"
)

func main() {

	cfgPath := flag.String("p", "./src/cmd/api/config.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Printf("Error while Loading config file\nReason: %v\n", err)
	}

	e := transport.New(cfg)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
