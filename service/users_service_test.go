package service_test

import (
	"testing"

	"boards.io/domain"
	"boards.io/service"
	"boards.io/transport/request"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UsersRepository struct {
	mock.Mock
}

func (r *UsersRepository) Create(user *domain.User) (string, error) {
	args := r.Called(user)
	return args.String(0), args.Error(1)
}

func makeUserRequestMock() request.NewUserReq {
	return request.NewUserReq{Name: faker.FirstName(), Username: faker.Username(), Password: faker.Password(), Email: faker.Email()}
}

func Test_Create_User(t *testing.T) {
	assert := assert.New(t)
	var userRequest = makeUserRequestMock()

	repositoryMock := new(UsersRepository)

	IDMock := "1"

	repositoryMock.On("Create", mock.Anything).Return(IDMock, nil)

	service := service.UsersService{
		Repository: repositoryMock,
	}

	ID, error := service.Create(userRequest)

	assert.Nil(error)
	assert.Equal(ID, IDMock)
}
