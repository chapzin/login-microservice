package domain

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int64      `json:"id" gorm:"primary_key;auto_increment"`
	Email        string     `json:"email" gorm:"unique_index;not null"`
	PasswordHash string     `json:"-"`
	ProfileImage string     `json:"avatar"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	TokenEmail   string     `json:"token"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func (u *User) Register(password string, confirm_pwd string, db *gorm.DB) (*User, error) {

	err := getEmail(db, u.Email)
	if err != nil {
		return nil, err
	}

	u.prepare()
	err = u.hashPassword(password, confirm_pwd)
	if err != nil {
		return nil, err
	}
	db.NewRecord(u)
	db.Create(u)
	return u, nil

}

func (u *User) prepare() {
	u.TokenEmail = uuid.NewV4().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

}

func (u *User) hashPassword(password, password_confirm string) error {
	if password != password_confirm {
		return fmt.Errorf("Senha diferente da confirmação")
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(pwdHash)
	return nil
}

func getEmail(db *gorm.DB, email string) error {
	var user User
	db.Where("email=?", email).First(&user)
	if user.ID != 0 {
		return fmt.Errorf("O Email %s já está cadastrado", email)
	}
	return nil
}
