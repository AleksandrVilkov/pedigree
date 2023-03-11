package storage

import "pedigree/internal/service/entiry"

type FamilyMemberDbInterface interface {
	DeleteFamilyMember(key string) error
	ReadFamilyMemberById(key string) (entiry.FamilyMember, error)
	UpdateFamilyMember(fm entiry.FamilyMember) entiry.FamilyMember
	CreateFamilyMembers(fm []entiry.FamilyMember) error
	ReadFamilyMembersByPedigreeUid(uid string) ([]entiry.FamilyMember, error)
}

type UserDBInterface interface {
	DeleteUser(key string) error
	UpdateUser(user entiry.User) entiry.User
	CreateUser(user entiry.User) error
	ReadUserById(key string) entiry.User
}
type PedigreeDBInterface interface {
	DeletePedigree(key string) error
	ReadPedigreeById(key string) (entiry.Pedigree, error)
	CreatePedigree(pedigree entiry.Pedigree) error
	UpdatePedigree(pedigree entiry.Pedigree) entiry.Pedigree
}
