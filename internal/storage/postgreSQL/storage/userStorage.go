package storage

import (
	"pedigree/internal/usecase"
)

type UserStorage struct {
	//Psql postgreSQL.PostgreSQL
}

func (us *UserStorage) CreateUser(user *usecase.User) (int, error) {
	//TODO
	return 0, nil
}
func (us *UserStorage) ReadUserById(username, password string) (usecase.User, error) {
	return usecase.User{}, nil
	//TODO
}
