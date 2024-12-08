package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}

func VerifyPassword(password string, input_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(input_password))
	return err == nil
}
