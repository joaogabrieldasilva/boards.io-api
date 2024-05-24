package domain

import (
	"errors"

	"boards.io/utils"
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID       string
	Name     string `validate:"required,min=5,max=30"`
	Email    string `validate:"required,email"`
	Username string `validate:"required,min=5"`
	Password string `validate:"required,min=6"`
}

func (user *User) Validate() error {
	validate := validator.New()

	if error := validate.Struct(user); error != nil {
		error := error.(validator.ValidationErrors)[0]
		return errors.New(utils.GetFieldErrorFromTag(error.Tag(), error.Field()))
	}

	return nil
}

type UsersRepository interface {
	Create(user *User) (string, error)
}
