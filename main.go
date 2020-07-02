package main

import (
	"fmt"
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

	user := domain.User{}

	userok, err := user.Register("12345678", "12345678", conn)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Usuario: %+v", userok)

}
