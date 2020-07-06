package repositories

import (
	"testing"

	"github.com/chapzin/login-microservice/framework"
	"github.com/stretchr/testify/require"
)

func TestNewUserRegister(t *testing.T) {
	db := framework.NewDbTest()
	defer db.Close()

	repo := NewUserRepositoryDb(db)
	user, err := repo.Register("chapzin@gmail.com", "123456", "123456")

	require.NotNil(t, user)
	require.Nil(t, err)
	require.Equal(t, "chapzin@gmail.com", user.Email)

}

func TestNewUserInsertPassword(t *testing.T) {
	db := framework.NewDbTest()
	defer db.Close()

	repo := NewUserRepositoryDb(db)
	user, err := repo.Register("teste@teste.com.br", "123456", "12345")
	require.Error(t, err)
	require.Nil(t, user)
}

func TestGetUserByEmail(t *testing.T) {
	db := framework.NewDbTest()
	defer db.Close()

	repo := NewUserRepositoryDb(db)
	user, err := repo.Register("chapzin@gmail.com", "123456", "123456")

	require.Nil(t, err)
	require.NotNil(t, user)

	getUser, err := repo.GetUserByEmail("chapzin@gmail.com")
	require.NotNil(t, getUser)
	require.Nil(t, err)
	require.Equal(t, getUser.Email, "chapzin@gmail.com")
	require.NotNil(t, getUser.ID)

}

func TestEmailRegistered(t *testing.T) {
	db := framework.NewDbTest()
	defer db.Close()

	repo := NewUserRepositoryDb(db)
	user, err := repo.Register("chapzin@gmail.com", "123456", "123456")

	require.Nil(t, err)
	require.NotNil(t, user)

	user2, err := repo.Register("chapzin@gmail.com", "456789", "456789")
	require.Error(t, err)
	require.Nil(t, user2)
}

func TestGetUsers(t *testing.T) {
	db := framework.NewDbTest()
	defer db.Close()

	repo := NewUserRepositoryDb(db)
	user, err := repo.Register("chapzin@gmail.com", "123456", "123456")
	require.Nil(t, err)
	require.NotNil(t, user)

	user2, err := repo.Register("chapzin2@gmail.com", "456789", "456789")
	require.Nil(t, err)
	require.NotNil(t, user2)
	cont := 0
	users := repo.GetUsers()
	require.NotNil(t, users)

	for _, u := range users {
		if u.ID != 0 {
			cont++
		}
	}

	require.Equal(t, 2, cont)

}
