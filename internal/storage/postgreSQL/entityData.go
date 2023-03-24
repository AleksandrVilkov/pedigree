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

func toFamilyMemberData(fm *usecase.FamilyMember) *FamilyMemberData {
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
	Name            string
	CreatedDate     string
	LastUpdatedDate string
	OwnerID         string
}

func (pd *PedigreeData) toModel() *usecase.Pedigree {
	createDate, _ := time.Parse(TEAMPLEATE_TIME, pd.CreatedDate)
	updateDate, _ := time.Parse(TEAMPLEATE_TIME, pd.LastUpdatedDate)
	return &usecase.Pedigree{
		ID:              "",
		Name:            pd.Name,
		CreatedDate:     createDate,
		LastUpdatedDate: updateDate,
		OwnerUid:        pd.OwnerID,
		PartOwners:      nil,
		FamilyMembers:   nil,
	}
}

func toPedigreeData(p *usecase.Pedigree) (*PedigreeData, error) {
	id, err := strconv.Atoi(p.ID)
	if err != nil {
		return nil, err
	}
	return &PedigreeData{
		ID:              id,
		Name:            p.Name,
		CreatedDate:     p.CreatedDate.Format(TEAMPLEATE_TIME),
		LastUpdatedDate: p.LastUpdatedDate.Format(TEAMPLEATE_TIME),
		OwnerID:         p.OwnerUid,
	}, nil
}

type UserData struct {
	ID              int
	Login           string
	CreatedDate     string
	LastUpdatedDate string
	Role            string
	FirstName       string
	LastName        string
	Password        string
	HasPedigree     string
}

func (u *UserData) ToModel() *usecase.User {
	createDate, _ := time.Parse(TEAMPLEATE_TIME, u.CreatedDate)
	updateDate, _ := time.Parse(TEAMPLEATE_TIME, u.LastUpdatedDate)
	return &usecase.User{
		ID:                   strconv.Itoa(u.ID),
		CreatedDate:          createDate,
		LastUpdatedDate:      updateDate,
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
		Login:           u.Login,
		CreatedDate:     u.CreatedDate.Format(TEAMPLEATE_TIME),
		LastUpdatedDate: u.CreatedDate.Format(TEAMPLEATE_TIME),
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
