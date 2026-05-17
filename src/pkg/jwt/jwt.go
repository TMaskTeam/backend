package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func InitJWTSecret(secret string) {
	secretKey = []byte(secret)
}

type Claims struct {
	OwnerID int    `json:"owner_id"`
	Role    string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(ownerID int, role string) (string, time.Time, error) {
	expiresAt := time.Now().Add(24 * time.Hour)

	claims := Claims{
		OwnerID: ownerID,
		Role:    role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

func ValidateToken(tokenString string) (int, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.OwnerID, claims.Role, nil
	}

	return 0, "", errors.New("invalid token")
}
