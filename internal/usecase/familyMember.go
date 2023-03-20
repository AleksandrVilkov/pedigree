package usecase

type FamilyMemberUsecase struct {
	Repository FamilyRepositoryInterface
}

func NewFamilyMemberUsecase(fm FamilyRepositoryInterface) *FamilyMemberUsecase {
	return &FamilyMemberUsecase{
		Repository: fm,
	}
}
func (f *FamilyMemberUsecase) Save()     {}
func (f *FamilyMemberUsecase) Delete()   {}
func (f *FamilyMemberUsecase) Update()   {}
func (f *FamilyMemberUsecase) Create()   {}
func (f *FamilyMemberUsecase) FindByID() {}
