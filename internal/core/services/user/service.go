package userService

import (
	"errors"
	"example/context/hello/internal/core/domain"
	"example/context/hello/internal/core/ports"
)

type service struct {
	usersRepository ports.UserRepository
}

func New(usersRepository ports.UserRepository,) *service {
	return &service{
		usersRepository: usersRepository,
	}
}


func (srv *service) Get(id string) (domain.User, error) {
	game, err := srv.usersRepository.Get(id)
	if err != nil {
		return domain.User{}, errors.New("get user from repository has failed")
	}

	return game, nil
}

func (srv *service) Create(name string, surname string, age string) (domain.User, error) {
	if name == "" && surname == "" && age == "" {
		return domain.User{}, errors.New("Error")
	}

	user := domain.User{Name: name, Surname: surname, Age: age}

	if err := srv.usersRepository.Save(user); err != nil {
		return domain.User{}, errors.New("create a new user error")
	}

	return user, nil
}