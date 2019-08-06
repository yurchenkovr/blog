package main

import (
	"blog/src/infrastructure/config"
	"blog/src/infrastructure/secure"
	"blog/src/models"
	"blog/src/repository/postgres"
	"flag"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"strings"
)

func main() {
	cfgPath := flag.String("p", "./src/cmd/api/config.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Printf("Error while Loading config file\nReason: %v\n", err)
	}

	db := postgres.New(cfg.DB.User, cfg.DB.Password, cfg.DB.Database, cfg.DB.Addr)

	CreateTables(db)
	InsertData(db)
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

func InsertData(db *pg.DB) {
	dbInsert := `INSERT INTO public.roles VALUES (200,200,'ADMIN');
				INSERT INTO public.roles VALUES (100,100,'USER');`
	queries := strings.Split(dbInsert, ";")

	for _, v := range queries[0 : len(queries)-1] {
		_, err := db.Exec(v)
		if err != nil {
			log.Printf("error while InsertData into roles\nReason: %v\n", err)
		}
	}

	sec := secure.HashAndSalt([]byte("admin"))

	adminInsert := `INSERT INTO public.users (id, username, password, created_at, role_id, blocked)
		VALUES (1, 'admin', '%s', now(), 200, false);`
	_, err := db.Exec(fmt.Sprintf(adminInsert, sec))
	if err != nil {
		log.Printf("error while InsertAdmin\nReason: %v\n", err)
	}
}
