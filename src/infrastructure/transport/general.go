package transport

import (
	"github.com/labstack/echo"
	"strconv"
)

func getParamID(c echo.Context) int {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		return 0
	}
	return id
}
