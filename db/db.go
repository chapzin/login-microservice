package db

import (
	"fmt"
	"log"
	"os"

	"github.com/chapzin/login-microservice/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	dbURL := os.Getenv("DB_URL")
	userDB := os.Getenv("USER_DB")
	pwdDB := os.Getenv("PWD_DB")
	nameDB := os.Getenv("NAME_DB")
	hostDB := os.Getenv("HOST_DB")
	if dbURL == "" {
		dbURL = fmt.Sprintf("%s:%s@(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", userDB, pwdDB, hostDB, nameDB)
	}
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		log.Fatal("Erro ao conectar ao mysql", err)
	}

	db.LogMode(true)

	db.AutoMigrate(&domain.User{})

	return db
}
