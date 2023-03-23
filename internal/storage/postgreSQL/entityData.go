package postgreSQL

import (
	"pedigree/internal/usecase"
	"strconv"
	"time"
)

type Message struct {
	Type MessageType
	Text string
}

type MessageType string

const (
	ERROR   MessageType = "ERROR"
	INFO                = "INFO"
	WARNING             = "WARNING"
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
	Login           string
	CreatedDate     time.Time
	LastUpdatedDate time.Time
	Role            string
	FirstName       string
	LastName        string
	Password        string
	HasPedigree     string
}

func (u *UserData) toModel() *usecase.User {
	return &usecase.User{
		ID:                   strconv.Itoa(u.ID),
		CreatedDate:          u.CreatedDate,
		LastUpdatedDate:      u.LastUpdatedDate,
		Role:                 usecase.Role(u.Role),
		FirstName:            u.FirstName,
		LastName:             u.LastName,
		Login:                u.Login,
		Password:             []byte(u.Password),
		CreatedPedigreesList: nil,
	}
}

func ToUserData(u *usecase.User) (*UserData, error) {
	res := &UserData{
		CreatedDate:     u.CreatedDate,
		LastUpdatedDate: u.LastUpdatedDate,
		Role:            u.Role.String(),
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Password:        string(u.Password),
	}

	if u.ID != "0" && u.ID != "" {
		idStr, err := strconv.Atoi(u.ID)
		if err != nil {
			return nil, err
		}
		res.ID = idStr
	}

	if len(u.CreatedPedigreesList) > 0 {
		res.HasPedigree = "true"
	} else {
		res.HasPedigree = "false"
	}

	return res, nil
}
