package storage

import (
	"pedigree/internal/db/postgreSQL"
	"pedigree/internal/service"
)

type FamalyMemberStorage struct {
	psql postgreSQL.PostgreSQL
}

func (fms *FamalyMemberStorage) DeleteFamilyMember(key string) error {
	return nil
}
func (fms *FamalyMemberStorage) ReadFamilyMemberById(key string) (service.FamilyMember, error) {
	return service.FamilyMember{}, nil
}
func (fms *FamalyMemberStorage) UpdateFamilyMember(fm service.FamilyMember) service.FamilyMember {
	return service.FamilyMember{}
}
func (fms *FamalyMemberStorage) CreateFamilyMembers(fm []service.FamilyMember) (int, error) {
	return 0, nil
}
func (fms *FamalyMemberStorage) ReadFamilyMembersByPedigreeUid(uid string) ([]service.FamilyMember, error) {
	return nil, nil
}
