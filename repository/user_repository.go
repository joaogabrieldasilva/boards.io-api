package repository

import "boards.io/domain"


type UsersRepository interface {
	Save(user *domain.User) error
}