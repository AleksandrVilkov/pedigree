package storage

import "pedigree/internal/usecase"

func FindRoleByString(strRole string) usecase.Role {
	switch strRole {
	case "ADMIN":
		return usecase.Role(1)
	case "USER":
		return usecase.Role(2)
	case "GUEST":
		return usecase.Role(3)
	default:
		return "UNKNOWN"
	}
}
