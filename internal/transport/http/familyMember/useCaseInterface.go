package familyMember

import (
	"pedigree/internal/usecase"
)

type FmUsecaseInterface interface {
	Save(u *usecase.FamilyMember) (int, error)
	Delete()
	Update()
	Create()
	FindByID()
}
