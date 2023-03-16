package postgreSQL

import (
	"pedigree/internal/usecase"
	"time"
)

type FamilyMemberData struct {
	ID                 int
	PedigreeID         int
	CreatedDate        time.Time
	LastUpdatedDate    time.Time
	FirstName          string
	LastName           string
	MiddleName         string
	MaidenName         string
	Dob                time.Time
	Dod                time.Time
	MainPlaceResidence string
	MotherID           int
	FatherID           int
	HasChildren        bool
	HasPartners        bool
	Note               string
}

func (fmd *FamilyMemberData) toModel() *usecase.FamilyMember {
	//TODO
	return &usecase.FamilyMember{
		ID:                 "",
		CreatedDate:        time.Time{},
		LastUpdatedDate:    time.Time{},
		PedigreeUid:        "",
		FirstName:          "",
		LastName:           "",
		MiddleName:         "",
		MaidenName:         "",
		DoB:                time.Time{},
		DoD:                time.Time{},
		MainPlaceResidence: usecase.Location{},
		Mother:             nil,
		Father:             nil,
		Children:           nil,
		Partners:           nil,
		Note:               "",
	}
}

func toPostgreSQLFamilyMember(fm *usecase.FamilyMember) *FamilyMemberData {
	return &FamilyMemberData{
		ID:                 0,
		PedigreeID:         0,
		CreatedDate:        time.Time{},
		LastUpdatedDate:    time.Time{},
		FirstName:          "",
		LastName:           "",
		MiddleName:         "",
		MaidenName:         "",
		Dob:                time.Time{},
		Dod:                time.Time{},
		MainPlaceResidence: "",
		MotherID:           0,
		FatherID:           0,
		HasChildren:        false,
		HasPartners:        false,
		Note:               "",
	}
}

type PedigreeData struct {
	ID              int
	CreatedDate     time.Time
	LastUpdatedDate time.Time
	OwnerID         int
}

func (pd *PedigreeData) toModel() *usecase.Pedigree {
	return &usecase.Pedigree{
		ID:              "",
		CreatedDate:     time.Time{},
		LastUpdatedDate: time.Time{},
		OwnerUid:        "",
		PartOwners:      nil,
		FamilyMembers:   nil,
	}
}

func toPedigreeData(p *usecase.Pedigree) *PedigreeData {
	return &PedigreeData{
		ID:              0,
		CreatedDate:     time.Time{},
		LastUpdatedDate: time.Time{},
		OwnerID:         0,
	}
}

type UserData struct {
	ID              int
	CreatedDate     time.Time
	LastUpdatedDate time.Time
	Role            string
	FirstName       string
	LastName        string
	Password        []byte
	HasPedigree     bool
}

func (u *UserData) toModel() *usecase.User {
	return &usecase.User{
		ID:                   "",
		CreatedDate:          time.Time{},
		LastUpdatedDate:      time.Time{},
		Role:                 0,
		FirstName:            "",
		LastName:             "",
		Password:             nil,
		CreatedPedigreesList: nil,
	}
}

func toUserData(u *usecase.User) *UserData {
	return &UserData{
		ID:              0,
		CreatedDate:     time.Time{},
		LastUpdatedDate: time.Time{},
		Role:            "",
		FirstName:       "",
		LastName:        "",
		Password:        nil,
		HasPedigree:     false,
	}
}
