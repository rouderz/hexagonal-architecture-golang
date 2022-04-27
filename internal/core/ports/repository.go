package ports

import (
	"example/context/hello/internal/core/domain"
)

type UserRepository interface {
	Get(id string) (domain.User, error)
	Save(domain.User) error
}

