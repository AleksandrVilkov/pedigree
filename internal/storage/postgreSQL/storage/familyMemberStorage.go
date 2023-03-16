package storage

import (
	"pedigree/internal/storage/postgreSQL"
	"pedigree/internal/usecase"
)

type FamalyMemberStorage struct {
	psql postgreSQL.PostgreSQL
}

func (fms *FamalyMemberStorage) DeleteFamilyMember(key string) error {
	return nil
}
func (fms *FamalyMemberStorage) ReadFamilyMemberById(key string) (usecase.FamilyMember, error) {
	return usecase.FamilyMember{}, nil
}
func (fms *FamalyMemberStorage) UpdateFamilyMember(fm usecase.FamilyMember) usecase.FamilyMember {
	return usecase.FamilyMember{}
}
func (fms *FamalyMemberStorage) CreateFamilyMembers(fm []usecase.FamilyMember) (int, error) {
	return 0, nil
}
func (fms *FamalyMemberStorage) ReadFamilyMembersByPedigreeUid(uid string) ([]usecase.FamilyMember, error) {
	return nil, nil
}
