package service

import "context"

type FamilyRepository interface {
	DeleteFamilyMember(key string) error
	ReadFamilyMemberById(key string) (FamilyMember, error)
	UpdateFamilyMember(fm FamilyMember) FamilyMember
	CreateFamilyMembers(fm []FamilyMember) (int, error)
	ReadFamilyMembersByPedigreeUid(uid string) ([]FamilyMember, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (int, error)
	ReadUserById(ctx context.Context, username, password string) (User, error)
}

type PedigreeRepository interface {
	DeletePedigree(key string) error
	ReadPedigreeById(key string) (Pedigree, error)
	CreatePedigree(pedigree Pedigree) (int, error)
	UpdatePedigree(pedigree Pedigree) Pedigree
}
