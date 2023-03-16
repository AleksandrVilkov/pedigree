package transport

import "context"

type AuthUseCaseInterface interface {
	SingUp(ctx context.Context, login, firstName, lastName, pass string) (int, error)
	SingIn(ctx context.Context, login, pass string) error
}
