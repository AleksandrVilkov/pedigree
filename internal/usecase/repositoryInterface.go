package usecase

import "context"

type FamilyRepositoryInterface interface {
	DeleteFamilyMember(key string) error
	ReadFamilyMemberById(key string) (FamilyMember, error)
	UpdateFamilyMember(fm FamilyMember) FamilyMember
	CreateFamilyMembers(fm []FamilyMember) (int, error)
	ReadFamilyMembersByPedigreeUid(uid string) ([]FamilyMember, error)
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *User) (int, error)
	ReadUserById(ctx context.Context, username, password string) (User, error)
}

type PedigreeRepositoryInterface interface {
	DeletePedigree(key string) error
	ReadPedigreeById(key string) (Pedigree, error)
	CreatePedigree(pedigree Pedigree) (int, error)
	UpdatePedigree(pedigree Pedigree) Pedigree
}
