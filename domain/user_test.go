package domain

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfUserIsEmpty(t *testing.T) {
	user := NewUser()
	err := user.Validate()

	require.Error(t, err)
}

func TestUserValidate(t *testing.T) {
	user := NewUser()
	user.Email = "chapzin@gmail.com"
	user.TokenEmail = uuid.NewV4().String()
	err := user.Validate()

	require.Nil(t, err)

}

func TestUserEmailIsNotValid(t *testing.T) {
	user := NewUser()
	user.Email = "avvvasdas"
	user.TokenEmail = uuid.NewV4().String()
	err := user.Validate()

	require.Error(t, err)
}

func TestUserUuidNotValid(t *testing.T) {
	user := NewUser()
	user.Email = "chapzin@gmail.com"
	user.TokenEmail = "abc"
	err := user.Validate()

	require.Error(t, err)
}
