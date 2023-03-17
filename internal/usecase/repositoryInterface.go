package usecase

type FamilyRepositoryInterface interface {
	DeleteFamilyMember(key string) error
	ReadFamilyMemberById(key string) (FamilyMember, error)
	UpdateFamilyMember(fm FamilyMember) FamilyMember
	CreateFamilyMembers(fm []FamilyMember) (int, error)
	ReadFamilyMembersByPedigreeUid(uid string) ([]FamilyMember, error)
}

type UserRepositoryInterface interface {
	CreateUser(user *User) (int, error)
	ReadUserById(username, password string) (User, error)
}

type PedigreeRepositoryInterface interface {
	DeletePedigree(key string) error
	ReadPedigreeById(key string) (Pedigree, error)
	CreatePedigree(pedigree Pedigree) (int, error)
	UpdatePedigree(pedigree Pedigree) Pedigree
}
