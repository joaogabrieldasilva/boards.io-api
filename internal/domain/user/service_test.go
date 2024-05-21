package user_test

import (
	"errors"
	"testing"

	"boards.io/internal/commonerrors"
	"boards.io/internal/domain/user"
	"boards.io/internal/dto"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(user *user.User) error {
	args := r.Called(user)
	return args.Error(0)
}

func makeUserDtoMock() dto.NewUserDto {
	return dto.NewUserDto{Name: faker.FirstName(), Username: faker.Username(), Password: faker.Password(), Email: faker.Email()}
}


func Test_Create_User(t *testing.T) {
	var userDto = makeUserDtoMock()

	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(user *user.User) bool {
		if userDto.Name != user.Name {
			return false
		} else if userDto.Username != user.Username {
			return false
		} else if userDto.Password != user.Password {
			return false
		} else if userDto.Email != user.Email {
			return false
		}

		return true
	})).Return(nil)

	service := user.Service{
		Repository: repositoryMock,
	}

	service.Create(userDto)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateUser(t *testing.T) {
	assert := assert.New(t)
	var userDto = makeUserDtoMock()
	userDto.Name = ""

	service := user.Service{}

	_, error := service.Create(userDto)

	assert.NotNil(error)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	var userDto = makeUserDtoMock()

	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.Anything).Return(commonerrors.ErrInternal)

	service := user.Service{
		Repository: repositoryMock,
	}

	_, error := service.Create(userDto)

	assert.True(errors.Is(commonerrors.ErrInternal, error))
}