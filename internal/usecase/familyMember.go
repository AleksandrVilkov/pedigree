package usecase

type FamilyMemberUsecase struct {
	Repository IFamilyRepository
}

func NewFamilyMemberUsecase(fm IFamilyRepository) *FamilyMemberUsecase {
	return &FamilyMemberUsecase{
		Repository: fm,
	}
}
func (f *FamilyMemberUsecase) Save(u *FamilyMember) (int, error) {
	// прекрутить базу данных, проверять на уже существующего пользователя
	return 0, nil
}
func (f *FamilyMemberUsecase) Delete()   {}
func (f *FamilyMemberUsecase) Update()   {}
func (f *FamilyMemberUsecase) Create()   {}
func (f *FamilyMemberUsecase) FindByID() {}
