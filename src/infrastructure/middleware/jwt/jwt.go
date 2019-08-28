package jwt

import (
	"blog/src/infrastructure/config"
	"blog/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Service struct {
	key            []byte
	expirationTime time.Duration
	algo           jwt.SigningMethod
}

func New(cfg *config.APIms) *Service {
	signingMethod := jwt.GetSigningMethod(cfg.JWT.SigningAlgorithm)
	if signingMethod == nil {
		panic("invalid jwt signing method")
	}

	duration, err := strconv.Atoi(cfg.JWT.Duration)
	if err != nil {
		log.Printf("Error when parsing duration: %v", err)
		return nil
	}
	return &Service{
		key:            []byte(cfg.JWT.Secret),
		algo:           signingMethod,
		expirationTime: time.Duration(duration) * time.Minute,
	}
}

func (s *Service) MWFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("token")
			if err != nil {
				if err == http.ErrNoCookie {
					log.Printf("cookie are absent ")
					return c.JSON(http.StatusUnauthorized, err)
				}
				return c.JSON(http.StatusBadRequest, err)
			}

			token, err := s.parseToken(cookie.Value)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, err)
			}

			claims := token.Claims.(jwt.MapClaims)

			id := int(claims["id"].(float64))
			username := claims["u"].(string)
			role := models.AccessRole(claims["r"].(float64))
			blocked := claims["b"].(bool)

			c.Set("id", id)
			c.Set("username", username)
			c.Set("isBlocked", blocked)
			c.Set("role", role)

			return next(c)
		}
	}
}

func (s *Service) parseToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(tkn *jwt.Token) (interface{}, error) {
		return s.key, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *Service) GenerateToken(u models.User) (string, error) {
	expire := time.Now().Add(s.expirationTime)

	token := jwt.NewWithClaims(s.algo, jwt.MapClaims{
		"id":  u.ID,
		"u":   u.Username,
		"b":   u.Blocked,
		"r":   u.Role.AccessLevel,
		"exp": expire.Unix(),
	})

	tokenString, err := token.SignedString(s.key)
	if err != nil {
		log.Println("error while generating token")
		return "", err
	}

	return tokenString, nil
}
