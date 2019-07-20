package repository

import (
	"blog/src/repository/controllers"
	"github.com/labstack/echo"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", controllers.MainCookie)
}
