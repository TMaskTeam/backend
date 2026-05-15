package password

import (
	"errors"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type TestHashPasswordContext struct {
	password      string
	expectedError error
}

type TestCheckPasswordContext struct {
	password       string
	hashedPassword string
	expectedError  error
}

func TestCheckPassword(t *testing.T) {
	Tests := []TestCheckPasswordContext{
		{
			password:       "password123",
			hashedPassword: "$2a$10$ZJuHJdEa0mzVXFzzK.91E.pfIU8L/4gkTKhQqFTE2YxzCSRTbDUfG",
			expectedError:  nil,
		},
		{
			password:       "password123",
			hashedPassword: "",
			expectedError:  ErrEmptyHashedPassword,
		},
		{
			password:       "password123",
			hashedPassword: "$2a$10$ZJuHJdEa0mzVXFzzK.91E.pfIU8L/4gkTKhQqFTE2YxzCSRTbDUfM",
			expectedError:  bcrypt.ErrMismatchedHashAndPassword,
		},
		{
			password:       "",
			hashedPassword: "$2a$10$ZJuHJdEa0mzVXFzzK.91E.pfIU8L/4gkTKhQqFTE2YxzCSRTbDUfG",
			expectedError:  ErrEmptyPassword,
		},
		{
			password:       "1",
			hashedPassword: "$2a$10$ZJuHJdEa0mzVXFzzK.91E.pfIU8L/4gkTKhQqFTE2YxzCSRTbDUfG",
			expectedError:  ErrInvalidPasswordLength,
		},
	}

	for _, test := range Tests {
		err := CheckHash(test.hashedPassword, test.password)
		if !errors.Is(err, test.expectedError) {
			t.Errorf("expected %v, got: %v", test.expectedError, err)
		}
	}
}

func TestHashPassword(t *testing.T) {
	Tests := []TestHashPasswordContext{
		{
			password:      "",
			expectedError: ErrEmptyPassword,
		},
		{
			password:      "1",
			expectedError: ErrInvalidPasswordLength,
		},
		{
			password:      "password",
			expectedError: nil,
		},
	}

	for _, test := range Tests {
		_, err := Hash(test.password)
		if !errors.Is(err, test.expectedError) {
			t.Errorf("expected %v, got: %v", test.expectedError, err)
		}
	}
}
