package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrEmptyPassword error = errors.New("Password can not be empty!")
var ErrEmptyHashedPassword error = errors.New("Hashed password can not be empty!")
var ErrInvalidPasswordLength error = errors.New("Length of password can be from 8 to 30!")

func Hash(password string) (string, error) {
	passwordLen := len(password)

	if passwordLen == 0 {
		return "", ErrEmptyPassword
	}

	if passwordLen < 8 {
		return "", ErrInvalidPasswordLength
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckHash(hashedPassword, password string) error {
	passwordLen := len(password)
	hashedPasswordLen := len(hashedPassword)

	if passwordLen == 0 {
		return ErrEmptyPassword
	}

	if passwordLen < 8 || passwordLen > 30 {
		return ErrInvalidPasswordLength
	}

	if hashedPasswordLen == 0 {
		return ErrEmptyHashedPassword
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
