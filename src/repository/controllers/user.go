package controllers

import (
	"blog/src/models"
	"blog/src/usecases"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
)

type serviceUser struct {
	svc usecases.UserService
}

func NewUserService(e echo.Echo, userService usecases.UserService) {
	userHTTPsvc := serviceUser{svc: userService}

	g := e.Group("/u")
	auth := g.Group("/log", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("my_secret_key"),
		TokenLookup: "cookie:token",
	}))
	adm := auth.Group("/admin")

	g.GET("/name/:username", userHTTPsvc.GetByUsername) //Any user can see 'user'
	g.GET("/:id", userHTTPsvc.GetByIDUser)              //
	g.GET("", userHTTPsvc.GetAllUsers)                  //
	g.POST("", userHTTPsvc.CreateUser)                  //

	auth.DELETE("/:id", userHTTPsvc.DeleteMyUser) //Auth user can Delete only his user profile
	auth.PATCH("/:id", userHTTPsvc.UpdateMyUser)  //				  Update

	adm.DELETE("/:id", userHTTPsvc.DeleteAnyUser)  // Admin can Delete any user profile
	adm.PATCH("/bl/:id", userHTTPsvc.BlockUser)    //			Block
	adm.PATCH("/unb/:id", userHTTPsvc.UnblockUser) //			Unblock
}
func (s *serviceUser) GetByUsername(c echo.Context) error {
	username := c.Param("username")

	user, err := s.svc.GetByUsername(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)
}
func (s *serviceUser) DeleteUser(c echo.Context) error {
	id := getParamID(c)
	err := s.svc.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Deleted")
}
func (s *serviceUser) GetByIDUser(c echo.Context) error {
	id := getParamID(c)
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
	id := getParamID(c)
	err = s.svc.UpdateUser(id, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "We got your updated user")
}
func (s *serviceUser) CreateUser(c echo.Context) error {
	user := models.User{}

	defer c.Request().Body.Close()

	b, errRead := ioutil.ReadAll(c.Request().Body)
	if errRead != nil {
		log.Printf("Fail reading the request body: %v", errRead)
		return c.JSON(http.StatusInternalServerError, errRead)
	}
	errUnmarsh := json.Unmarshal(b, &user)
	if errUnmarsh != nil {
		log.Printf("Fail unmarshaling in CraateUser: %v", errUnmarsh)
		return c.JSON(http.StatusInternalServerError, errUnmarsh)
	}
	req := usecases.CreateReqUser{
		Username: user.Username,
		Password: user.Password,
		RoleID:   user.RoleID,
	}
	err := s.svc.SaveUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "we got your user!")
}
func (s *serviceUser) GetAllUsers(c echo.Context) error {
	users, err := s.svc.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, users)
}
func (s *serviceUser) UpdateMyUser(c echo.Context) error {
	if mine := s.isMine(c); mine != true {
		log.Printf("Error: It`s not your article\n")
		return c.String(http.StatusBadRequest, "Please choose your article")
	}
	if err := s.UpdateUser(c); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, "")
}
func (s *serviceUser) DeleteMyUser(c echo.Context) error {
	if mine := s.isMine(c); mine != true {
		log.Printf("Error: It`s not your article\n")
		return c.String(http.StatusBadRequest, "Please choose your article")
	}
	if err := s.DeleteUser(c); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, "Deleted successfully")
}
func (s *serviceUser) DeleteAnyUser(c echo.Context) error {
	claims, err := getClaims(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if isAdmin(c, claims.ID) == true {
		if err := s.DeleteUser(c); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.String(http.StatusOK, "Deleted successfully")
	}
	return c.JSON(http.StatusBadRequest, err)
}
func (s *serviceUser) BlockUser(c echo.Context) error {
	claims, err := getClaims(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	id := getParamID(c)
	if isAdmin(c, claims.ID) == true {
		if err := s.svc.BlockUser(id); err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return c.String(http.StatusOK, "User Blocked.")
	}
	return c.String(http.StatusBadRequest, "You`re not an ADMIN")
}

func (s *serviceUser) UnblockUser(c echo.Context) error {
	claims, err := getClaims(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	id := getParamID(c)
	if isAdmin(c, claims.ID) == true {
		if err := s.svc.UnblockUser(id); err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return c.String(http.StatusOK, "User Unblocked.")
	}
	return c.String(http.StatusBadRequest, "You`re not an ADMIN")
}
func (s *serviceUser) isMine(c echo.Context) bool {
	id := getParamID(c)
	claims, errClaims := getClaims(c)
	if errClaims != nil {
		return false
	}
	art, err := s.svc.GetByIDUser(id)
	if err != nil {
		return false
	}
	if art.Username != claims.Username {
		log.Printf("Error: Please choose your article\n")
		return false
	}
	return true
}
