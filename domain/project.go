package domain


type Project struct {
	ID string
	Name string
	Users []User
}

func CreateProject(name string, users []User) *Project {
	return &Project{
		ID: "1",
		Name: name,
		Users: users,
	}
}