package usecase

import (
	"crypto/sha1"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	pkgJwt "pedigree/pkg/jwt"
	"time"
)

type AuthUsecase struct {
	userRepository IUserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUsecase(userRepository IUserRepository,
	hashSalt string,
	signingKey []byte,
	expireDuration time.Duration) *AuthUsecase {
	return &AuthUsecase{
		userRepository: userRepository,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

func (a *AuthUsecase) SingUp(login, firstName, lastName, pass string) (int, error) {
	user := &User{
		CreatedDate:          time.Now(),
		LastUpdatedDate:      time.Now(),
		Role:                 USER,
		Login:                login,
		FirstName:            firstName,
		LastName:             lastName,
		Password:             []byte(pass),
		CreatedPedigreesList: nil,
	}
	return a.userRepository.CreateUser(user)
}

func (a *AuthUsecase) SingIn(login, pass string) (string, error) {
	user, err := a.userRepository.ReadUserByUserName(login)

	if pass != string(user.Password) {
		return "", errors.New("Invalid password")
	}

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &pkgJwt.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
	})
	return token.SignedString(a.signingKey)
}

func getPassWithSalt(pass, salt string) string {
	pwd := sha1.New()
	pwd.Write([]byte(pass))
	pwd.Write([]byte(salt))
	return string(pwd.Sum(nil))
}
