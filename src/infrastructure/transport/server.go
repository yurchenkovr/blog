package transport

import (
	"blog/src/infrastructure/config"
	"blog/src/infrastructure/middleware/jwt"
	"blog/src/repository/postgres"
	"blog/src/usecases"
	"blog/src/usecases/rbac"
	"github.com/labstack/echo"
	"github.com/nats-io/nats.go"
	"log"
)

func New(cfg *config.APIms) *echo.Echo {
	e := echo.New()

	dbHandler := postgres.New(cfg)

	jwtService := jwt.New(cfg)

	rbac := rbac.Service{}
	nc, err := nats.Connect(cfg.Nats.Url)
	if err != nil {
		log.Fatalf("error when connect to nats")
	}

	NewService(*e, usecases.NewArtService(postgres.NewArticleRepository(dbHandler), &rbac, nc), jwtService.MWFunc())
	NewUserService(*e, usecases.NewUserService(postgres.NewUserRepository(dbHandler), jwtService, &rbac, nc), jwtService.MWFunc())

	return e
}
