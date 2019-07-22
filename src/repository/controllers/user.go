package controllers

import (
	"blog/src/models"
	"blog/src/usecases"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

type serviceUser struct {
	svc usecases.UserService
}

func NewUserService(e echo.Echo, userService usecases.UserService) {
	userHTTPsvc := serviceUser{svc: userService}

	g := e.Group("/users")

	g.GET("", userHTTPsvc.GetAllUsers)
	g.POST("", userHTTPsvc.CreateUser)
	g.DELETE("/:id", userHTTPsvc.DeleteUser)
	g.PATCH("/:id", userHTTPsvc.UpdateUser)
	g.GET("/:id", userHTTPsvc.GetByIDUser)
}
func (s *serviceUser) DeleteUser(c echo.Context) error {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		return errID
	}
	err := s.svc.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Deleted successfully")
}
func (s *serviceUser) GetByIDUser(c echo.Context) error {
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		return errID
	}
	user, err := s.svc.GetByIDUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (s *serviceUser) UpdateUser(c echo.Context) error {
	user := models.User{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Fail reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Printf("Fail unmarshaling in updateUser: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	id, errID := strconv.Atoi(c.Param("id"))
	if errID != nil {
		return errID
	}

	err = s.svc.UpdateUser(id, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Printf("This is your user: %#v", user)
	return c.JSON(http.StatusOK, "We got your updated user")
}
func (s *serviceUser) CreateUser(c echo.Context) error {
	user := models.User{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Fail reading the request body: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Printf("Fail unmarshaling in CraateUser: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	req := usecases.CreateReqUser{
		Username: user.Username,
		Password: user.Password,
		RoleID:   user.RoleID,
	}
	err = s.svc.SaveUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Printf("This is your user: %v", user)
	return c.JSON(http.StatusOK, "we got your user!")
}
func (s *serviceUser) GetAllUsers(c echo.Context) error {
	users, err := s.svc.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, users)
}
