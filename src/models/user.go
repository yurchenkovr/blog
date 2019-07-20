package models

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"-"`
	Role     *Role  `json:"role,omitempty"`

	RoleID AccessRole `json:"-"`
}
