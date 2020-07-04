package repositories

import (
	"testing"

	"github.com/chapzin/login-microservice/framework"
	"github.com/stretchr/testify/require"
)

func TestNewUserRepositoryDBInsert(t *testing.T) {
	db := framework.NewDbTest()
	defer db.Close()

	repo := NewUserRepositoryDb(db)
	user, err := repo.Register("chapzin@gmail.com", "123456", "123456")

	require.NotEmpty(t, user.ID)
	require.Nil(t, err)
	require.Equal(t, user.Email, "chapzin@gmail.com")

}

func TestNewUserRepositoryDBInsertPassword(t *testing.T) {
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
