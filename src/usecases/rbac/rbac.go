package rbac

import (
	"blog/src/models"
	"github.com/labstack/echo"
)

// RBAC represents role-based-access-control interface
type RBAC interface {
	EnforceUser(echo.Context, int) bool
	IsAdmin(echo.Context) bool
	IsBlocked(echo.Context) bool
}

// New creates new RBAC service
func New() *Service {
	return &Service{}
}

// Service is RBAC application service
type Service struct {
}

func (s *Service) IsAdmin(c echo.Context) bool {
	return c.Get("role").(models.AccessRole) == models.AdminRole
}

// EnforceUser checks whether the request to change user data is done by the same user
func (s *Service) EnforceUser(c echo.Context, ID int) bool {
	return (c.Get("id") == ID) || s.IsAdmin(c)
}

func (s Service) IsBlocked(c echo.Context) bool {
	return c.Get("isBlocked").(bool)
}
