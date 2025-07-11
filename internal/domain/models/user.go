// This package conatins the domain models for the chat application.
package models

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        string `json:"id" validate:"uuid4"`
	Username  string `json:"username" validate:"required,min=3,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,sha256"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
