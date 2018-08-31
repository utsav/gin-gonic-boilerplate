package services

import "github.com/utsav/gin-gonic-boilerplate/structs"

func TodoData() structs.Todo {
	return structs.Todo{"title", "description"}
}