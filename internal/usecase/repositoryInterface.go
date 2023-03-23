package usecase

type IFamilyRepository interface {
	DeleteFamilyMember(key string) error
	ReadFamilyMemberById(key string) (FamilyMember, error)
	UpdateFamilyMember(fm FamilyMember) FamilyMember
	CreateFamilyMembers(fm []FamilyMember) (int, error)
	ReadFamilyMembersByPedigreeUid(uid string) ([]FamilyMember, error)
}

type IUserRepository interface {
	CreateUser(user *User) (int, error)
	ReadUserByUserName(username string) (User, error)
}

type IPedigreeRepository interface {
	DeletePedigree(key string) error
	ReadPedigreeById(key string) (Pedigree, error)
	CreatePedigree(pedigree Pedigree) (int, error)
	UpdatePedigree(pedigree Pedigree) Pedigree
}
