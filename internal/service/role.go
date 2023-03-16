package service

type Role int

const (
	ADMIN   Role = 1
	USER         = 2
	GUEST        = 3
	UNKNOWN      = 4
)

func (r Role) String() string {
	switch r {
	case ADMIN:
		return "ADMIN"
	case USER:
		return "USER"
	case GUEST:
		return "GUEST"
	default:
		return "UNKNOWN"
	}
}
