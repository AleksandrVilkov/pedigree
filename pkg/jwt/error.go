package jwt

import "errors"

var ErrInvalidAccessToken = errors.New("invalid auth token")
var ErrUserDoesNotExist = errors.New("user does not exist")
var ErrUserAlreadyExists = errors.New("user with such credentials already exist")
var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
