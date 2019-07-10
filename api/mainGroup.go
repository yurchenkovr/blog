package api

import (
	"github.com/labstack/echo"
	"blog/api/handlers"
)

func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.Yallo)
	e.GET("/login", handlers.Login)
}
