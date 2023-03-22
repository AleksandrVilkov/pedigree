package usecase

type Role string

const (
	ADMIN   Role = "ADMIN"
	USER         = "USER"
	GUEST        = "GUEST"
	UNKNOWN      = "UNKNOWN"
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
