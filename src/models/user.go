package models

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
	Role     *Role  `json:"role,omitempty"`

	RoleID AccessRole `json:"role_id"`
}
