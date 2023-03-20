package storage

import (
	"pedigree/internal/usecase"
)

type FamilyMemberStorage struct {
	//psql postgreSQL.PostgreSQL
}

func (fms *FamilyMemberStorage) DeleteFamilyMember(key string) error {
	return nil
}
func (fms *FamilyMemberStorage) ReadFamilyMemberById(key string) (usecase.FamilyMember, error) {
	return usecase.FamilyMember{}, nil
}
func (fms *FamilyMemberStorage) UpdateFamilyMember(fm usecase.FamilyMember) usecase.FamilyMember {
	return usecase.FamilyMember{}
}
func (fms *FamilyMemberStorage) CreateFamilyMembers(fm []usecase.FamilyMember) (int, error) {
	return 0, nil
}
func (fms *FamilyMemberStorage) ReadFamilyMembersByPedigreeUid(uid string) ([]usecase.FamilyMember, error) {
	return nil, nil
}
