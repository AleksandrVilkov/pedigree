package pedigree

type pedigreeUsecaseInterface interface {
	Save()
	Delete()
	Update()
	Create()
	FindByID()
}
