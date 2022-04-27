package userRepositories

import (
	"encoding/json"
	"errors"
	"example/context/hello/internal/core/domain"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (domain.User, error) {
	if value, ok := repo.kvs[id]; ok {
		user := domain.User{}
		err := json.Unmarshal(value, &user)
		if err != nil {
			return domain.User{}, errors.New("fail to get value from kvs")
		}

		return user, nil
	}

	return domain.User{}, errors.New("user not found in kvs")
}

func (repo *memkvs) Save(User domain.User) error {
	bytes, err := json.Marshal(User)
	if err != nil {
		return errors.New("user fails at marshal into json string")
	}

	repo.kvs[User.ID] = bytes

	return nil
}