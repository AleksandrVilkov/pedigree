package familyMember

import (
	"pedigree/internal/usecase"
)

type IFmUsecase interface {
	Save(u *usecase.FamilyMember) (int, error)
	Delete()
	Update()
	Create()
	FindByID()
}
