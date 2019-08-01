package transport

import (
	"blog/src/infrastructure/middleware/jwt"
	"blog/src/repository/postgres"
	"blog/src/usecases"
	"blog/src/usecases/rbac"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	dbHandler := postgres.New()
	jwtService := jwt.New("my_secret_key", "HS256", 5)
	rbac := rbac.Service{}

	NewService(*e, usecases.NewArtService(postgres.NewArticleRepository(dbHandler), &rbac), jwtService.MWFunc())
	NewUserService(*e, usecases.NewUserService(postgres.NewUserRepository(dbHandler), jwtService, &rbac), jwtService.MWFunc())

	return e
}
