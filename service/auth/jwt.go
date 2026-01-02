package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kasariks/notes_api/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userID":    userID,
		"expiresAt": config.Envs.JWTExpirationInSeconds,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
