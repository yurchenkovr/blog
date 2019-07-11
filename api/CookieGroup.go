package api

import (
	"blog/api/handlers"
	"github.com/labstack/echo"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)
}
