package test

import (
	"github.com/google/uuid"
	"pedigree/internal/db/storage"
	"pedigree/internal/service/entiry"
	"strconv"
	"testing"
	"time"
)

func TestCreatePedigree(t *testing.T) {
	psql := storage.Storage{}
	count := 3
	familyMembers := make([]entiry.FamilyMember, count)
	p := entiry.Pedigree{
		Uid:             uuid.New().String(),
		CreatedDate:     time.Now(),
		LastUpdatedDate: time.Now(),
		FamilyMembers:   familyMembers,
	}

	for i := 0; i < count; i++ {
		familyMembers[i] = entiry.FamilyMember{
			Uid:             uuid.New().String(),
			CreatedDate:     time.Now(),
			LastUpdatedDate: time.Now(),
			PedigreeUid:     p.Uid,
			LastName:        "last name test " + strconv.Itoa(i),
			FirstName:       "first name test " + strconv.Itoa(i),
			DoB:             time.Now(),
		}
	}

	err := psql.CreatePedigree(p)
	if err != nil {
		t.Fatal()
	}
}
