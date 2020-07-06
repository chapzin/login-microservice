package main

import (
	"fmt"
	"log"

	"github.com/chapzin/login-microservice/application/repositories"
	"github.com/chapzin/login-microservice/framework"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao tenta abrir .env")
	}
}

func main() {
	db := framework.NewDbProduction()
	defer db.Close()
	repo := repositories.NewUserRepositoryDb(db)
	user, err := repo.Register("chapzin@gmail.com", "123456", "123456")

	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", user)
}
