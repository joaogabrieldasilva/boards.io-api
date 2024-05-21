package project

import (
	"boards.io/internal/domain/user"
)

type Project struct {
	ID string
	Name string
	Users []user.User
}

func CreateProject(name string, users []user.User) *Project {
	return &Project{
		ID: "1",
		Name: name,
		Users: users,
	}
}