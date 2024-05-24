package repository

import (
	"database/sql"
	"fmt"

	"boards.io/domain"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		db,
	}
}

func (repository *UsersRepository) Create(user *domain.User) (string, error) {

	statement, error := repository.db.Prepare("INSERT INTO users (name, email, username, password)")

	if error != nil {
		return "", domain.ErrInternal
	}

	result, error := statement.Exec(user.Name, user.Email, user.Username, user.Password)

	if error != nil {
		return "", domain.ErrInternal
	}

	ID, error := result.LastInsertId()

	if error != nil {
		return "", domain.ErrInternal
	}

	return fmt.Sprint(ID), nil
}
