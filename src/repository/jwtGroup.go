package repository

import (
	"blog/src/repository/controllers"
	"github.com/labstack/echo"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", controllers.MainJwt)
}
