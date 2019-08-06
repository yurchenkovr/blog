package postgres

import (
	"blog/src/models"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/gommon/log"
)

func New() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "vitmantekoR_2408",
		Database: "simpleblog",
		Addr:     "localhost:5432",
	})
	if db == nil {
		log.Printf("Failed to connect to database!\n")
	}
	fmt.Printf("Connection to database successful.\n")

	CreateTables(db)

	return db
}

func CreateTables(db *pg.DB) {
	for _, model := range []interface{}{&models.User{}, &models.Article{}, &models.Role{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Printf("Error while creating table: %v\nReason: %v\n", model, err)
		}
	}
}
