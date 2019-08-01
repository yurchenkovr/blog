package transport

import (
	"blog/src/models"
	"blog/src/usecases"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
)

type serviceUser struct {
	svc usecases.UserService
}

func NewUserService(e echo.Echo, userService usecases.UserService, middlewareFunc echo.MiddlewareFunc) {
	userHTTPsvc := serviceUser{svc: userService}

	ur := e.Group("/users")

	ur.POST("/login", userHTTPsvc.Login)

	ur.GET("/name/:username", userHTTPsvc.GetByUsername)
	ur.GET("/:id", userHTTPsvc.View)
	ur.GET("", userHTTPsvc.List)
	ur.POST("", userHTTPsvc.Create)
	ur.DELETE("/:id", userHTTPsvc.Delete, middlewareFunc)
	ur.PATCH("/:id", userHTTPsvc.Update, middlewareFunc)
	ur.PATCH("/bl/:id", userHTTPsvc.Block, middlewareFunc)
	ur.PATCH("/unb/:id", userHTTPsvc.Unblock, middlewareFunc)
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *serviceUser) Login(c echo.Context) error {
	loginReq := loginRequest{}
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Fail reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &loginReq)
	if err != nil {
		log.Printf("Fail unmarshaling in login request: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	token, err := s.svc.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		log.Printf("Fail unmarshaling in login request: %s", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Path = "/"
	c.SetCookie(cookie)
	return nil
}

func (s *serviceUser) GetByUsername(c echo.Context) error {
	username := c.Param("username")

	user, err := s.svc.GetByUsername(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (s *serviceUser) Delete(c echo.Context) error {
	id := getParamID(c)

	err := s.svc.Delete(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "Deleted")
}

func (s *serviceUser) View(c echo.Context) error {
	id := getParamID(c)

	user, err := s.svc.View(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (s *serviceUser) Update(c echo.Context) error {
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

	err = s.svc.Update(c, id, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "We got your updated user")
}

func (s *serviceUser) Create(c echo.Context) error {
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

	err := s.svc.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "we got your user!")
}

func (s *serviceUser) List(c echo.Context) error {
	users, err := s.svc.List()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (s *serviceUser) Block(c echo.Context) error {
	id := getParamID(c)

	if err := s.svc.Block(c, id); err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.String(http.StatusOK, "User Blocked.")
}

func (s *serviceUser) Unblock(c echo.Context) error {
	id := getParamID(c)

	if err := s.svc.Unblock(c, id); err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.String(http.StatusOK, "User Unblocked.")
}
