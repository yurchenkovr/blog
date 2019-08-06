package transport

import (
	"blog/src/infrastructure/config"
	"blog/src/infrastructure/middleware/jwt"
	"blog/src/repository/postgres"
	"blog/src/usecases"
	"blog/src/usecases/rbac"
	"github.com/labstack/echo"
)

func New(cfg *config.Configuration) *echo.Echo {
	e := echo.New()

	dbHandler := postgres.New(cfg.DB.User, cfg.DB.Password, cfg.DB.Database, cfg.DB.Addr)

	jwtService := jwt.New(cfg.JWT.Secret, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)

	rbac := rbac.Service{}

	NewService(*e, usecases.NewArtService(postgres.NewArticleRepository(dbHandler), &rbac), jwtService.MWFunc())
	NewUserService(*e, usecases.NewUserService(postgres.NewUserRepository(dbHandler), jwtService, &rbac), jwtService.MWFunc())

	return e
}
