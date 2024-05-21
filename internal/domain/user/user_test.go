package user_test

import (
	"fmt"
	"testing"

	"boards.io/internal/domain/user"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)


func makeUserMock() (name, username, email, password string) {
	name = faker.Name()
	username = faker.Username()
	email = "joao@gmail.com"
	password = faker.Password()

	fmt.Println(email)

	return
}

func Test_User_CreateUser(t *testing.T) {
	assert := assert.New(t)
	
	name, username, email, password := makeUserMock()

	user, error := user.CreateUser(name,username, email, password)

	assert.Equal(error, nil)
	assert.NotNil(user.ID)
	assert.Equal(user.Name, name)
	assert.Equal(user.Username, username)
	assert.Equal(user.Password, password)
}

func Test_User_ValidateName(t *testing.T) {
	assert := assert.New(t)
	
	_, username, email, password := makeUserMock()
	_, error := user.CreateUser("", username, email, password)

	assert.Equal(error.Error(), "Name is required")
}

func Test_User_ValidateUsername(t *testing.T) {
	assert := assert.New(t)
	
	name, _, email, password := makeUserMock()
	_, error := user.CreateUser(name, "", email, password)

	assert.Equal(error.Error(), "Username is required")
}

func Test_User_ValidateEmail(t *testing.T) {
	assert := assert.New(t)
	
	name, username, _, password := makeUserMock()
	_, error := user.CreateUser(name, username, "", password)

	assert.Equal(error.Error(), "Email is required")
}

func Test_User_ValidatePassword(t *testing.T) {
	assert := assert.New(t)
	
	name, username, email, _ := makeUserMock()
	_, error := user.CreateUser(name, username, email, "")

	assert.Equal(error.Error(), "Password is required")
}