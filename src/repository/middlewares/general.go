package middlewares

import "github.com/labstack/echo"

func SetMainMiddleares(e *echo.Echo) {
	e.Use(ServerHeader)
}
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set(echo.HeaderServer, "Blog/1.0")
		return next(c)
	}
}
