package main

import (
	"fmt"

	"boards.io/internal/domain/user"
)


func main() {


	if _, error := user.CreateUser("", "", "", ""); error !=nil {
		fmt.Println(error.Error())
	}
}