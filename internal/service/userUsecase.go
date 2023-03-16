package service

import "context"

type UserUsecase interface {
	SingUp(ctx context.Context) error
	SingIn(ctx context.Context) (string, error)
	ParseToken(ctx context.Context) (User, error)
}
