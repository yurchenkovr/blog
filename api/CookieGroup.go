package api

import (
	"github.com/labstack/echo"
	"blog/api/handlers"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)
}
