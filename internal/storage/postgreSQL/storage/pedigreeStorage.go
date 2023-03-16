package storage

import (
	"pedigree/internal/storage/postgreSQL"
	"pedigree/internal/usecase"
)

type PedigreeStorage struct {
	psql postgreSQL.PostgreSQL
}

func (ps *PedigreeStorage) DeletePedigree(key string) error {
	return nil
}
func (ps *PedigreeStorage) ReadPedigreeById(key string) (usecase.Pedigree, error) {
	return usecase.Pedigree{}, nil
}
func (ps *PedigreeStorage) CreatePedigree(pedigree usecase.Pedigree) (int, error) {
	return 0, nil
}
func (ps *PedigreeStorage) UpdatePedigree(pedigree usecase.Pedigree) usecase.Pedigree {
	return usecase.Pedigree{}
}
