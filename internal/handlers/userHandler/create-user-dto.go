package userHandler

import "example/context/hello/internal/core/domain"

type BodyCreate struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     string `json:"age"`
}

type ResponseCreate domain.User

func BuildResponseCreate(model domain.User) ResponseCreate {
	return ResponseCreate(model)
}
