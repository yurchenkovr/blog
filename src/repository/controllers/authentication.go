package controllers

import (
	"blog/src/infrastructure/secure"
	"blog/src/models"
	"blog/src/usecases"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var jwtKey = []byte("my_secret_key")

type service struct {
	svc usecases.UserService
}

func NewAuthService(e echo.Echo, userServ usecases.UserService) {
	authHTTPsvc := service{svc: userServ}

	e.POST("/signin", authHTTPsvc.Signin)
	e.GET("/welcome", authHTTPsvc.Welcome)
	e.POST("/refresh", authHTTPsvc.Refresh)
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
type Claims struct {
	Username string            `json:"username"`
	ID       int               `json:"id"`
	Role     models.AccessRole `json:"role"`
	Blocked  bool              `json:"blocked"`
	jwt.StandardClaims
}

func (s *service) Signin(c echo.Context) error {
	var creds Credentials

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Fail reading the request body: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = json.Unmarshal(b, &creds)
	if err != nil {
		log.Printf("Fail unmarshaling creds: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	user, err := s.svc.GetByUsername(creds.Username)
	if err != nil {
		log.Printf("error: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	if truePass := secure.ComparePasswords(user.Password, []byte(creds.Password)); !truePass {
		return c.JSON(http.StatusUnauthorized, err)
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		ID:       user.ID,
		Role:     user.RoleID,
		Blocked:  user.Blocked,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = expirationTime
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "Sign in successful")
}

func (a *service) Welcome(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Printf("cookie are absent ")
			return c.JSON(http.StatusUnauthorized, err)
		}
		return c.JSON(http.StatusBadRequest, err)
	}

	tokenString := cookie.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return c.JSON(http.StatusBadRequest, err)
	}
	if !token.Valid {
		return c.JSON(http.StatusUnauthorized, err)
	}
	return c.String(http.StatusOK, fmt.Sprintf("Welcome %s\nID: %d\nRole: %d\nExpires at: %v\nBlocked: %t\n", claims.Username, claims.ID, claims.Role, claims.ExpiresAt, claims.Blocked))
}
func (a *service) Refresh(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err != http.ErrNoCookie {
			return c.JSON(http.StatusUnauthorized, err)
		}
		fmt.Println("error:1")
		return c.JSON(http.StatusBadRequest, err)
	}
	tokenString := cookie.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		return c.JSON(http.StatusUnauthorized, err)
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, err)
		}
		fmt.Println("error:2")
		return c.JSON(http.StatusBadRequest, err)
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		fmt.Println("error:3")
		return c.JSON(http.StatusBadRequest, err)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tknStr, err := tkn.SignedString(jwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	newCookie := new(http.Cookie)
	newCookie.Name = "token"
	newCookie.Value = tknStr
	newCookie.Expires = expirationTime
	c.SetCookie(newCookie)

	return c.String(http.StatusOK, "token was refreshed")
}

func getClaims(c echo.Context) (*Claims, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Printf("cookie are absent ")
			return nil, c.JSON(http.StatusUnauthorized, err)
		}
		fmt.Println("something bad")
		return nil, c.JSON(http.StatusBadRequest, err)
	}
	tokenString := cookie.Value
	claims := &Claims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return claims, nil
}
