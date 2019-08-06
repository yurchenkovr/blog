package postgres

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/labstack/gommon/log"
)

func New(user, password, database, addr string) *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     addr,
	})
	if db == nil {
		log.Printf("Failed to connect to database!\n")
	}
	fmt.Printf("Connection to database successful.\n")

	return db
}
