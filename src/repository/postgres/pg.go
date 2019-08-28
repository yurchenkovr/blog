package postgres

import (
	"blog/src/infrastructure/config"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/labstack/gommon/log"
)

func New(cfg *config.APIms) *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Database: cfg.DB.Database,
		Addr:     cfg.DB.Addr,
	})
	if db == nil {
		log.Printf("Failed to connect to database!\n")
	}
	fmt.Printf("Connection to database successful.\n")

	return db
}
