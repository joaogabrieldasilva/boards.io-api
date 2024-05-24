package service

import (
	"boards.io/domain"
	"boards.io/repository"
	"boards.io/transport/request"
)



type UsersService struct {
	Repository repository.UsersRepository
}

func (s *UsersService) Create(userDto request.NewUserReq) (string, error) {

	user, error := domain.CreateUser(userDto.Name, userDto.Username, userDto.Email,userDto.Password)

	if error != nil {
		return "", error
	}

	error = s.Repository.Save(user)

	if error != nil {
		return "", domain.ErrInternal
	}

	return user.ID, nil
} 