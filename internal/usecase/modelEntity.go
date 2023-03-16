package usecase

import (
	"time"
)

type User struct {
	ID                   string
	CreatedDate          time.Time
	LastUpdatedDate      time.Time
	Role                 Role
	Login                string
	FirstName            string
	LastName             string
	Password             []byte
	CreatedPedigreesList []string
}

type Pedigree struct {
	ID              string
	CreatedDate     time.Time
	LastUpdatedDate time.Time
	OwnerUid        string
	PartOwners      []User
	FamilyMembers   []FamilyMember
}

type FamilyMember struct {
	ID                 string
	CreatedDate        time.Time
	LastUpdatedDate    time.Time
	PedigreeUid        string
	FirstName          string
	LastName           string
	MiddleName         string
	MaidenName         string
	DoB                time.Time
	DoD                time.Time
	MainPlaceResidence Location
	Mother             *FamilyMember
	Father             *FamilyMember
	Children           []FamilyMember
	Partners           []FamilyMember
	Note               string
}

type Location struct {
	Country Country
	City    City
}

type Country struct {
	Name string
}

type City struct {
	Name    string
	country Country
}
