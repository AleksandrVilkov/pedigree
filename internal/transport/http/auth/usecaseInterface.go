package auth

type IAuthUsecase interface {
	SingUp(login, firstName, lastName, pass string) (int, error)
	SingIn(login, pass string) (string, error)
}
