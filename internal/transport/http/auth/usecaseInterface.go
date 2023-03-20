package auth

type AuthUsecaseInterface interface {
	SingUp(login, firstName, lastName, pass string) (int, error)
	SingIn(login, pass string) (string, error)
}
