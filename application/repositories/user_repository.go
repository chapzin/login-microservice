package repositories

import (
	"fmt"
	"time"

	"github.com/chapzin/login-microservice/domain"
	"github.com/chapzin/login-microservice/framework"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	Register(email string, password string, password_cofirm string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetUsers() []domain.User
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func NewUserRepositoryDb(db *gorm.DB) *UserRepositoryDb {
	return &UserRepositoryDb{Db: db}
}

func (userRepo *UserRepositoryDb) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := userRepo.Db.Where("email=?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *UserRepositoryDb) Register(email string, password string, confirm_pwd string) (*domain.User, error) {

	user, err := userRepo.GetUserByEmail(email)
	if user != nil {
		return nil, fmt.Errorf("Email já cadastrado")
	}

	passHash, err := framework.HashPassword(password, confirm_pwd)

	if err != nil {
		return nil, err
	}

	user = domain.NewUser()
	user.Email = email
	user.TokenEmail = uuid.NewV4().String()
	user.PasswordHash = passHash
	user.CreatedAt = time.Now()

	userRepo.Db.NewRecord(user)
	userRepo.Db.Create(user)
	return user, nil

}

func (userRepo *UserRepositoryDb) GetUsers() []domain.User {
	var users []domain.User
	userRepo.Db.Order("id asc").Find(&users)
	return users
}
