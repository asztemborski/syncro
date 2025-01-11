package model_test

import (
	"testing"

	"github.com/asztemborski/syncro/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	account := model.NewAccount("test@example.com", "test_username")

	assert.Equal(t, account.Email, "test@example.com")
	assert.Equal(t, account.Username, "test_username")
	assert.Equal(t, account.EmailVerified, false)
	assert.Equal(t, account.IsActive, true)
}

func TestAccountNormalizeEmail(t *testing.T) {
	account := model.Account{Email: " Test@example.com"}
	account.NormalizeEmail()

	assert.Equal(t, account.Email, "test@example.com")
}

func TestAccountNormalizeUsername(t *testing.T) {
	account := model.Account{Username: "Test_Username "}
	account.NormalizeUsername()

	assert.Equal(t, account.Username, "test_username")
}

func TestPasswordSet(t *testing.T) {
	account := model.NewAccount("test@example.com", "test_username")
	account.Password.Set("secret_password")

	assert.True(t, account.Password.Compare("secret_password"))
}
