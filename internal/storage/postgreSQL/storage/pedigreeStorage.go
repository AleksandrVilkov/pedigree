package storage

import (
	"pedigree/internal/storage/postgreSQL"
	"pedigree/internal/usecase"
)

type PedigreeStorage struct {
	psql postgreSQL.PostgreSQL
}

const (
	PEDIGREE_TABLE_NAME   = "PEDIGREE "
	P_NAME_COLUMN         = "name"
	P_ID_COLUMN           = "id"
	P_CREATED_DATE_COLUMN = "createdDate"
	P_LAST_UPDATED_COLUMN = "lastUpdatedDate"
	P_OWNER_UID_COLUMN    = "ownerUid"
)

func (ps *PedigreeStorage) DeletePedigree(key string) error {
	return nil
}
func (ps *PedigreeStorage) ReadPedigreeById(key string) (usecase.Pedigree, error) {
	return usecase.Pedigree{}, nil
}

func (ps *PedigreeStorage) CreatePedigree(pedigree usecase.Pedigree) (int, error) {

	return 0, nil
}
func (ps *PedigreeStorage) createPedigreeQuery(pg *postgreSQL.PedigreeData) string {
	return postgreSQL.INSERT_INTO + PEDIGREE_TABLE_NAME +
		"(" + P_NAME_COLUMN + ", " +
		P_CREATED_DATE_COLUMN + ", " +
		P_LAST_UPDATED_COLUMN + ", " +
		P_OWNER_UID_COLUMN + "), " +
		postgreSQL.VALUES + "(" +
		" '" + pg.Name + "', " +
		" '" + pg.CreatedDate + "', " +
		" '" + pg.LastUpdatedDate + "', " +
		" '" + pg.OwnerID + "')"
}
func (ps *PedigreeStorage) UpdatePedigree(pedigree usecase.Pedigree) usecase.Pedigree {
	return usecase.Pedigree{}
}
