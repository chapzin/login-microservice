package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID           int64      `valid:"-" json:"id" gorm:"primary_key;auto_increment"`
	Email        string     `valid:"email,required" json:"email" gorm:"unique_index;notnull"`
	PasswordHash string     `valid:"-" json:"-"`
	ProfileImage string     `valid:"-" json:"avatar"`
	FirstName    string     `valid:"-" json:"first_name"`
	LastName     string     `valid:"-" json:"last_name"`
	TokenEmail   string     `valid:"required,uuid" json:"token"`
	CreatedAt    time.Time  `valid:"-" json:"created_at"`
	UpdatedAt    time.Time  `valid:"-" json:"updated_at"`
	DeletedAt    *time.Time `valid:"-" json:"deleted_at"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	return nil
}
