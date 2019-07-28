package models

type AccessRole int

const (
	AdminRole AccessRole = 200
	UserRole  AccessRole = 100
)

type Role struct {
	ID          AccessRole `json:"id"`
	AccessLevel AccessRole `json:"access_level"`
	Name        string     `json:"name"`
}
