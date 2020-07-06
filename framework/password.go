package framework

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, password_confirm string) (string, error) {
	if password != password_confirm {
		return "", fmt.Errorf("Senha diferente da confirmação")
	}
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwdHash), nil
}

func VerifyPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
