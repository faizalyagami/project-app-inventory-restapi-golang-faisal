package model

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
