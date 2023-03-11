package storage

import (
	"pedigree/internal/db"
	"pedigree/internal/db/entityData"
	"pedigree/internal/service/entiry"
	"strconv"
)

func (s *Storage) CreateFamilyMembers(fms []entiry.FamilyMember) error {
	connector := db.PostgreSQL{}
	_, err := connector.SendQuery(createFamilyMembersQuery(fms))
	if err != nil {
		return err
	}
	return nil
}

// TODO
func (s *Storage) DeleteFamilyMember(key string) error {
	return nil
}

// TODO получить уиды детей
func (s *Storage) ReadChildrenUids(familyMemberUid string) []string {
	return nil
}

// TODO получить уиды партнерв
func (s *Storage) ReadPartnerIids(familyMemberUid string) []string {
	return nil
}

func (s *Storage) ReadFamilyMemberById(key string) (entiry.FamilyMember, error) {
	query := db.SELECT + "* " + db.FROM + db.FAMILY_MEMBERS_TABLE_NAME + db.WHERE +
		" uid = '" + key + "';"

	connector := db.PostgreSQL{}
	row := connector.GetRow(query)
	var result entityData.FamilyMemberIndex
	err := row.Scan(&result.Uid,
		&result.PedigreeUid,
		&result.CreatedDate,
		&result.LastUpdatedDate,
		&result.FirstName,
		&result.LastName,
		&result.MiddleName,
		&result.MaidenName,
		&result.Dob,
		&result.Dod,
		&result.MainPlaceResidence,
		&result.MotherUid,
		&result.FatherUid,
		&result.HasChildren,
		&result.HasPartners,
		&result.Note)
	if err != nil {
		return entiry.FamilyMember{}, err
	}
	return fillFamilyMemberFromIndex(result), nil
}

func (s *Storage) ReadFamilyMembersByPedigreeUid(uid string) ([]entiry.FamilyMember, error) {
	query := db.SELECT + "* " + db.FROM + db.FAMILY_MEMBERS_TABLE_NAME + db.WHERE +
		" pedigreeUid = '" + uid + "';"

	connector := db.PostgreSQL{}
	rows := connector.GetRows(query)

	var result []entiry.FamilyMember
	for rows.Next() {
		var fm entityData.FamilyMemberIndex
		err := rows.Scan(&fm.Uid,
			&fm.PedigreeUid,
			&fm.CreatedDate,
			&fm.LastUpdatedDate,
			&fm.FirstName,
			&fm.LastName,
			&fm.MiddleName,
			&fm.MaidenName,
			&fm.Dob,
			&fm.Dod,
			&fm.MainPlaceResidence,
			&fm.MotherUid,
			&fm.FatherUid,
			&fm.HasChildren,
			&fm.HasPartners,
			&fm.Note)
		if err != nil {
			return nil, err
		}
		result = append(result, fillFamilyMemberFromIndex(fm))
	}
	return result, nil
}

// TODO
func (s *Storage) UpdateFamilyMember(fm entiry.FamilyMember) entiry.FamilyMember {
	return entiry.FamilyMember{}
}

// TODO
func createFamilyMembersQuery(fms []entiry.FamilyMember) string {

	query := db.INSERT_INTO + db.FAMILY_MEMBERS_TABLE_NAME + "(uid, pedigreeUid, createdDate, lastUpdatedDate, firstName, lastName, " +
		"middleName, maidenName, dob, dod, mainPlaceResidence, motherUid, fatherUid, hasChildren, hasPartners, note)" + db.VALUES
	for i := 0; len(fms) > i; i++ {
		fm := fms[i]
		hasChildren := false

		if len(fm.Children) > 0 {
			hasChildren = true
		}
		hasPartners := false

		if len(fm.Partners) > 0 {
			hasPartners = true
		}
		query += "('" + fm.Uid + "', " +
			"'" + fm.PedigreeUid + "', " +
			" '" + fm.CreatedDate.Format(db.DATE_PATTERN) + "'," +
			" '" + fm.LastUpdatedDate.Format(db.DATE_PATTERN) + "'," +
			"'" + fm.FirstName + "', " +
			"'" + fm.LastName + "', " +
			"'" + fm.MiddleName + "', " +
			"'" + fm.MaidenName + "', " +
			"'" + fm.DoB.Format(db.DATE_PATTERN) + "', " +
			"'" + fm.DoD.Format(db.DATE_PATTERN) + "', " +
			"'" + fm.MainPlaceResidence.City.Name + "', "

		if fm.Mother != nil {
			query += "'" + fm.Mother.Uid + "', "
		} else {
			query += "'null', "
		}

		if fm.Father != nil {
			query += "'" + fm.Father.Uid + "', "
		} else {
			query += "'null', "
		}

		query += "'" + strconv.FormatBool(hasChildren) + "', " +
			"'" + strconv.FormatBool(hasPartners) + "', " +
			"'" + fm.Note + "'"
		if i+1 == len(fms) {
			query += ");"
		} else {
			query += "), "
		}
	}
	return query
}
