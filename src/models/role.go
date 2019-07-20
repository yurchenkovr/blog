package models

type AccessRole int

const (
	AdminRole = 200
	UserRole  = 100
)

type Role struct {
	ID          AccessRole `json:"id"`
	AccessLevel AccessRole `json:"access_level"`
	Name        string     `json:"name"`
}
