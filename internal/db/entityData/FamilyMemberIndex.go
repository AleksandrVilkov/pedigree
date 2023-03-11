package entityData

import "time"

type FamilyMemberIndex struct {
	Uid                string
	PedigreeUid        string
	CreatedDate        time.Time
	LastUpdatedDate    time.Time
	FirstName          string
	LastName           string
	MiddleName         string
	MaidenName         string
	Dob                time.Time
	Dod                time.Time
	MainPlaceResidence string
	MotherUid          string
	FatherUid          string
	HasChildren        bool
	HasPartners        bool
	Note               string
}
