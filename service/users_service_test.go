package service_test

import (
	"errors"
	"testing"

	"boards.io/domain"
	"boards.io/service"
	"boards.io/transport/request"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(user *domain.User) error {
	args := r.Called(user)
	return args.Error(0)
}

func makeUserDtoMock() request.NewUserReq {
	return request.NewUserReq{Name: faker.FirstName(), Username: faker.Username(), Password: faker.Password(), Email: faker.Email()}
}


func Test_Create_User(t *testing.T) {
	var userDto = makeUserDtoMock()

	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(user *domain.User) bool {
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

	service := service.UsersService{
		Repository: repositoryMock,
	}

	service.Create(userDto)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateUser(t *testing.T) {
	assert := assert.New(t)
	var userDto = makeUserDtoMock()
	userDto.Name = ""

	service := service.UsersService{}

	_, error := service.Create(userDto)

	assert.NotNil(error)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	var userDto = makeUserDtoMock()

	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.Anything).Return(domain.ErrInternal)

	service := service.UsersService{
		Repository: repositoryMock,
	}

	_, error := service.Create(userDto)

	assert.True(errors.Is(domain.ErrInternal, error))
}