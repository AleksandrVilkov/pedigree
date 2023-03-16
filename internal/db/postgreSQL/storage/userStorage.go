package storage

import (
	"context"
	"pedigree/internal/db/postgreSQL"
	"pedigree/internal/service"
)

type UserStorage struct {
	psql postgreSQL.PostgreSQL
}

func (us *UserStorage) CreateUser(ctx context.Context, user service.User) (int, error) {
	//TODO
	return 0, nil
}
func (us *UserStorage) ReadUserById(ctx context.Context, username, password string) (service.User, error) {
	return service.User{}, nil
	//TODO
}
