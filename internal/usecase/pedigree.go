package usecase

type PedigreeUsecase struct {
	pedigreeStorage IPedigreeRepository
}

// TODO
func SavePedigree(pedigree Pedigree) error {
	return nil
}

// TODO
func ReadPedigree(uid string) Pedigree {
	return Pedigree{}
}

// TODO
func DeletePedigree(uid string) error {
	return nil
}

// TODO
func UpdatePedigree(uid string) Pedigree {
	return Pedigree{}
}
