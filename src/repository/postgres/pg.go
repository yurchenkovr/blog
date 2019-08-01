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

	CreateTableArticle(db)

	return db
}
func CreateTableArticle(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&models.Article{}, opts)
	if err != nil {
		log.Printf("Error while creating table Article, Reason: %v\n", err)
	}
}
