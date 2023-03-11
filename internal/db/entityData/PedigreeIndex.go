package entityData

import "time"

type PedigreeIndex struct {
	Uid             string
	CreatedDate     time.Time
	LastUpdatedDate time.Time
	OwnerUid        string
}
