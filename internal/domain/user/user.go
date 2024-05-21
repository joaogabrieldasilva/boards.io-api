package user

import (
	"errors"

	"boards.io/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)


type User struct {
	ID string `validate:"required"`
	Name string `validate:"required,min=5,max=30"`
	Email string `validate:"required,email"`
	Username string `validate:"required,min=5"`
	Password string `validate:"required,min=6"`
}


func CreateUser(name string,  username string, email string,password string) (*User, error) {

	newUser := &User{
		ID: xid.New().String(),
		Name: name,
		Email: email,
		Username: username,
		Password: password,
	}

	validate := validator.New()

	if error := validate.Struct(newUser); error != nil {
		error := error.(validator.ValidationErrors)[0]
		
		return nil, errors.New(validation.GetFieldErrorFromTag(error.Tag(), error.Field()))
	}

	return newUser, nil
}