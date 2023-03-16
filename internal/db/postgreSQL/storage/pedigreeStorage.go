package storage

import (
	"pedigree/internal/db/postgreSQL"
	"pedigree/internal/service"
)

type PedigreeStorage struct {
	psql postgreSQL.PostgreSQL
}

func (ps *PedigreeStorage) DeletePedigree(key string) error {
	return nil
}
func (ps *PedigreeStorage) ReadPedigreeById(key string) (service.Pedigree, error) {
	return service.Pedigree{}, nil
}
func (ps *PedigreeStorage) CreatePedigree(pedigree service.Pedigree) (int, error) {
	return 0, nil
}
func (ps *PedigreeStorage) UpdatePedigree(pedigree service.Pedigree) service.Pedigree {
	return service.Pedigree{}
}
