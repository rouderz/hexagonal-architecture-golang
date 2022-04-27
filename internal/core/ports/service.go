package ports

import "example/context/hello/internal/core/domain"

type UserService interface {
	Get(id string) (domain.User, error)
	Create(name string, surname string, age string) (domain.User, error)
}
