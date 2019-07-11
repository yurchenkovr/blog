package api

import (
	"blog/api/handlers"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.Yallo)
	e.GET("/login", handlers.Login)
}
