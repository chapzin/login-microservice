package auth

import (
	"fmt"

	"github.com/chapzin/login-microservice/application/repositories"
	"github.com/chapzin/login-microservice/framework"
)

func SignIn(email string, password string) (string, error) {
	// TODO: Corrigir o acesso a instancia de conexao, deve receber a instancia pelo contexto
	db := framework.NewDbTest()
	userRepo := repositories.NewUserRepositoryDb(db)
	user, err := userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("email ou senha está errado")
	}

	err = framework.VerifyPassword(user.PasswordHash, password)
	if err != nil {
		return "", fmt.Errorf("email ou senha está errado")
	}

	token, err := framework.GenerateJWT(*user)
	if err != nil {
		return "", err
	}
	return token, nil
}
