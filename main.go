package main

import (
	"log"

	"github.com/chapzin/login-microservice/db"
	"github.com/chapzin/login-microservice/domain"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao tenta abrir .env")
	}
}

func main() {

	conn := db.Connect()

	user := domain.User{
		Email:           "chapzin@gmail.com",
		Password:        "123456",
		PasswordConfirm: "123555",
	}

	user.Register(conn)

}
