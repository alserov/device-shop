package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckPassword(password string, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}
