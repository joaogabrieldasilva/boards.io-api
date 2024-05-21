package user

import (
	"boards.io/internal/commonerrors"
	"boards.io/internal/dto"
)



type Service struct {
	Repository Repository
}

func (s *Service) Create(userDto dto.NewUserDto) (string, error) {

	user, error := CreateUser(userDto.Name, userDto.Username, userDto.Email,userDto.Password)

	if error != nil {
		return "", error
	}

	error = s.Repository.Save(user)

	if error != nil {
		return "", commonerrors.ErrInternal
	}

	return user.ID, nil
} 