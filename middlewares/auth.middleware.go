package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UID uint `json:"uid"`
	jwt.RegisteredClaims
}

func Authenticate(userId uint, secret string) (string, error) {
	expiredAt := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	})

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}