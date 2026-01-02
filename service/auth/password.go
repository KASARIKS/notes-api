package auth

import (
	"crypto/sha512"
	"fmt"
)

var IncorrectPassword error = fmt.Errorf("incorrect password")

func HashPassword(password string) (string, error) {
	hasher := sha512.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}

	hash := hasher.Sum(nil)

	return string(hash), nil
}

func ComparePasswords(currentPassHashed, gottenPass string) error {
	hashed, err := HashPassword(gottenPass)
	if err != nil {
		return err
	}

	if currentPassHashed != hashed {
		return IncorrectPassword
	}

	return nil
}
