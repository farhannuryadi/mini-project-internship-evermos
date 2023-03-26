package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedByte), nil
}

func CheckPasswordHas(password string, userPassword string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	return err == nil
}