package main

import (
	"blog/src/infrastructure/config"
	"blog/src/infrastructure/secure"
	"blog/src/models"
	"blog/src/repository/postgres"
	gC "blog/src/usecases/grpc/client"
	"flag"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"strings"
)

func main() {
	grpcPath := flag.String("p", "./src/cmd/cmdmanager/grpcConfig.yaml", "Path to gRPC config file")
	flag.Parse()

	cfg, err := config.Load(*grpcPath)
	if err != nil {
		log.Printf("Error while Loading config file\nReason: %v\n", err)
	}
	config := gC.Configs(cfg.Grpc.Port)

	db := postgres.New(config.APIms)

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
	dbInsert := `INSERT INTO public.roles (id, access_level, name) 
					SELECT 200,200,'ADMIN' 
				 WHERE NOT EXISTS (
					SELECT id FROM public.roles WHERE id = 200);
				 INSERT INTO public.roles (id, access_level, name) 
					SELECT 100,100,'USER'
				 WHERE NOT EXISTS (
					SELECT id FROM public.roles WHERE id = 100);`
	queries := strings.Split(dbInsert, ";")

	for _, v := range queries[0 : len(queries)-1] {
		_, err := db.Exec(v)
		if err != nil {
			log.Printf("error while InsertData into roles\nReason: %v\n", err)
		}
	}

	sec := secure.HashAndSalt([]byte("admin"))

	adminInsert := `INSERT INTO public.users (id, username, password, created_at, role_id, blocked)
				 		SELECT 1, 'admin', '%s', now(), 200, false 
					WHERE NOT EXISTS (
						SELECT id FROM public.users WHERE username = 'admin');`
	_, err := db.Exec(fmt.Sprintf(adminInsert, sec))
	if err != nil {
		log.Printf("error while InsertAdmin\nReason: %v\n", err)
	}
}
