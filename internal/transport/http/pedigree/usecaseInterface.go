package pedigree

type IPedigreeUsecase interface {
	Save()
	Delete()
	Update()
	Create()
	FindByID()
}
