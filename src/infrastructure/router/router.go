package router

import (
	"blog/src/repository"
	"blog/src/repository/controllers"
	"blog/src/repository/middlewares"
	"blog/src/repository/postgres"
	"blog/src/usecases"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	dbHandler := postgres.New()

	controllers.NewService(*e, usecases.NewArtService(postgres.NewArticleRepository(dbHandler)))
	controllers.NewUserService(*e, usecases.NewUserService(postgres.NewUserRepository(dbHandler)))
	controllers.NewAuthService(*e, usecases.NewUserService(postgres.NewUserRepository(dbHandler)))

	//articleGroup := e.Group("/article")

	//set all middlewares
	middlewares.SetMainMiddleares(e)
	//middlewares.SetArticleMiddlewares(articleGroup)
	//middlewares.SetCookieMiddlewares(cookieGroup)
	//middlewares.SetJWTMiddlewares(jwtGroup)

	//set main routes
	repository.MainGroup(e)

	return e
}
