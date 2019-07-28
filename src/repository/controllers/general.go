package controllers

import (
	"blog/src/models"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the web side!")
}

func isAdmin(c echo.Context, claimsID int) bool {
	claims, err := getClaims(c)
	if err != nil {
		return false
	}
	if claims.Role == models.AdminRole {
		return true
	}
	fmt.Println("You`re not an admin")
	return false
}
func getParamID(c echo.Context) int {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		return 0
	}
	return id
} /*
func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		claims, err := getClaims(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if claims.Role == models.AdminRole {
			return c.JSON(http.StatusOK, nil)
		}
		fmt.Println("You`re not an admin")
		return c.JSON(http.StatusBadRequest, err)
	}
}*/
