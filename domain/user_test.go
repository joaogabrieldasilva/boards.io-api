package domain_test

import (
	"testing"

	"boards.io/domain"
	"github.com/go-faker/faker/v4"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func makeUserMock() domain.User {
	return domain.User{
		ID:       xid.New().String(),
		Name:     faker.Name(),
		Username: faker.Username(),
		Email:    "joao@gmail.com",
		Password: faker.Password(),
	}
}

func Test_User_Validate(t *testing.T) {
	assert := assert.New(t)

	user := makeUserMock()

	error := user.Validate()
	assert.Nil(error)

}

func Test_User_ValidateName(t *testing.T) {
	assert := assert.New(t)

	user := makeUserMock()
	user.Name = ""
	error := user.Validate()
	assert.Equal(error.Error(), "Name is required")
}

func Test_User_ValidateUsername(t *testing.T) {
	assert := assert.New(t)

	user := makeUserMock()
	user.Username = ""

	error := user.Validate()

	assert.Equal(error.Error(), "Username is required")
}

func Test_User_ValidateEmail(t *testing.T) {
	assert := assert.New(t)

	user := makeUserMock()
	user.Email = ""

	error := user.Validate()

	assert.Equal(error.Error(), "Email is required")
}

func Test_User_ValidatePassword(t *testing.T) {
	assert := assert.New(t)

	user := makeUserMock()
	user.Password = ""

	error := user.Validate()

	assert.Equal(error.Error(), "Password is required")
}
