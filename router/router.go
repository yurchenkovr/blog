package router

import (
	"github.com/labstack/echo"
	"blog/api"
	"blog/api/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	//adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")
	articleGroup := e.Group("/article")

	//set all middlewares
	middlewares.SetMainMiddleares(e)
	middlewares.SetArticleMiddlewares(articleGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJWTMiddlewares(jwtGroup)

	//set main routes
	api.MainGroup(e)

	//set group routes
	api.ArticleGroup(articleGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	return e
}
