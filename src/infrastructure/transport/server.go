package transport

import (
	"blog/src/infrastructure/config"
	"blog/src/infrastructure/middleware/jwt"
	"blog/src/repository/postgres"
	"blog/src/usecases"
	"blog/src/usecases/rbac"
	"github.com/labstack/echo"
)

func New(cfg *config.APIms) *echo.Echo {
	e := echo.New()

	dbHandler := postgres.New(cfg)

	jwtService := jwt.New(cfg)

	rbac := rbac.Service{}

	NewService(*e, usecases.NewArtService(postgres.NewArticleRepository(dbHandler), &rbac), jwtService.MWFunc())
	NewUserService(*e, usecases.NewUserService(postgres.NewUserRepository(dbHandler), jwtService, &rbac), jwtService.MWFunc())

	return e
}
