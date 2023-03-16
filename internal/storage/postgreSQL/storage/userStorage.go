package storage

import (
	"context"
	"pedigree/internal/usecase"
)

type UserStorage struct {
	//Psql postgreSQL.PostgreSQL
}

func (us *UserStorage) CreateUser(ctx context.Context, user *usecase.User) (int, error) {
	//TODO
	return 0, nil
}
func (us *UserStorage) ReadUserById(ctx context.Context, username, password string) (usecase.User, error) {
	return usecase.User{}, nil
	//TODO
}
