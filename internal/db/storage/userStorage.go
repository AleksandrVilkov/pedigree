package storage

import "pedigree/internal/service/entiry"

func (s *Storage) CreateUser(user entiry.User) error {
	return nil
}

func (s *Storage) DeleteUser(key string) error {
	return nil
}

func (s *Storage) UpdateUser(user entiry.User) entiry.User {
	return entiry.User{}
}

func (s *Storage) ReadUserById(key string) entiry.User {
	return entiry.User{}
}
