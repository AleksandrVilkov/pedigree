package usecase

import (
	"context"
	"crypto/sha1"
	"time"
)

type AuthUseCase struct {
	userRepository UserRepositoryInterface
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(userRepository UserRepositoryInterface,
	hashSalt string,
	signingKey []byte,
	expireDuration time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepository: userRepository,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

func (a *AuthUseCase) SingUp(ctx context.Context, login, firstName, lastName, pass string) (int, error) {
	pwd := sha1.New()
	pwd.Write([]byte(pass))
	pwd.Write([]byte(a.hashSalt))
	user := &User{
		CreatedDate:          time.Now(),
		LastUpdatedDate:      time.Now(),
		Role:                 USER,
		Login:                login,
		FirstName:            firstName,
		LastName:             lastName,
		Password:             pwd.Sum(nil),
		CreatedPedigreesList: nil,
	}
	return a.userRepository.CreateUser(ctx, user)
}

func (a *AuthUseCase) SingIn(ctx context.Context, login, pass string) error {
	//TODO ищем пользователя
	return nil
}
