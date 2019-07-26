package repository

import (
	"blog/src/repository/controllers"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/", controllers.Yallo)
}
