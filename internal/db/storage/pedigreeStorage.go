package storage

import (
	"pedigree/internal/db"
	"pedigree/internal/db/entityData"
	"pedigree/internal/service/entiry"
)

type Storage struct {
}

func (s *Storage) CreatePedigree(pedigree entiry.Pedigree) error {
	connector := db.PostgreSQL{}

	_, err := connector.SendQuery(createPedigreeQuery(&pedigree))
	if err != nil {
		return err
	}
	if len(pedigree.FamilyMembers) > 0 {
		err := s.CreateFamilyMembers(pedigree.FamilyMembers)

		if err != nil {
			return err
		}
	}

	if len(pedigree.PartOwners) > 0 {
		//TODO сохранять совладельцев родослвной, если они есть
	}

	return nil
}

func createPedigreeQuery(pedigree *entiry.Pedigree) string {
	return db.INSERT_INTO + db.PEDIGREE_TABLE_NAME + " (uid, createddate, lastupdateddate, owneruid) " + db.VALUES +
		"('" + pedigree.Uid + "', " +
		" '" + pedigree.CreatedDate.Format(db.DATE_PATTERN) + "'," +
		" '" + pedigree.LastUpdatedDate.Format(db.DATE_PATTERN) + "'," +
		" '" + pedigree.OwnerUid + "');"
}

// TODO
func (s *Storage) DeletePedigree(key string) error {
	return nil
}

func (s *Storage) ReadPedigreeById(key string) (entiry.Pedigree, error) {
	query := db.SELECT + "* " + db.FROM + db.PEDIGREE_TABLE_NAME + db.WHERE +
		" uid = '" + key + "';"

	connector := db.PostgreSQL{}
	row := connector.GetRow(query)
	var result entityData.PedigreeIndex
	err := row.Scan(&result.Uid,
		&result.CreatedDate,
		&result.LastUpdatedDate,
		&result.OwnerUid)
	if err != nil {
		return entiry.Pedigree{}, err
	}
	return fillPedigreeFromIndex(result), nil

}

// TODO
func (s *Storage) UpdatePedigree(pedigree entiry.Pedigree) entiry.Pedigree {
	return entiry.Pedigree{}
}
