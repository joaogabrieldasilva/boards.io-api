package service

import (
	"boards.io/domain"
	"boards.io/transport/request"
)

type UsersService struct {
	Repository domain.UsersRepository
}

func NewUsersService(usersRepository domain.UsersRepository) *UsersService {
	return &UsersService{
		Repository: usersRepository,
	}
}

func (service *UsersService) Create(userReq request.NewUserReq) (string, error) {

	user := &domain.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Username: userReq.Username,
		Password: userReq.Password,
	}

	error := user.Validate()

	if error != nil {
		return "", error
	}

	ID, error := service.Repository.Create(user)

	if error != nil {
		return "", domain.ErrInternal
	}

	return ID, nil
}
