package storage

import (
	"pedigree/internal/db/entityData"
	"pedigree/internal/service/entiry"
)

func fillFamilyMemberFromIndex(fi entityData.FamilyMemberIndex) entiry.FamilyMember {
	result := entiry.FamilyMember{
		Uid:             fi.Uid,
		CreatedDate:     fi.CreatedDate,
		LastUpdatedDate: fi.LastUpdatedDate,
		PedigreeUid:     fi.PedigreeUid,
		FirstName:       fi.FirstName,
		LastName:        fi.LastName,
		MiddleName:      fi.MiddleName,
		MaidenName:      fi.MaidenName,
		DoD:             fi.Dod,
		DoB:             fi.Dob,
		Note:            fi.Note,
	}

	//TODO получать родителей, детей партнеров
	//TODO получать локацию
	return result
}
func fillPedigreeFromIndex(pi entityData.PedigreeIndex) entiry.Pedigree {
	result := entiry.Pedigree{
		Uid:             pi.Uid,
		CreatedDate:     pi.CreatedDate,
		LastUpdatedDate: pi.LastUpdatedDate,
		OwnerUid:        pi.OwnerUid,
	}
	//TODO получать членов семьи
	return result
}
