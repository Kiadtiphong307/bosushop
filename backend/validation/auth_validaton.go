package validation

import (
	"github.com/go-playground/validator/v10"
)

type RegisterInput struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}


type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

var validate = validator.New()

func ValidateRegisterInput(input RegisterInput) error {
	return validate.Struct(input)
}

func ValidateLoginInput(input LoginInput) error {
	return validate.Struct(input)
}