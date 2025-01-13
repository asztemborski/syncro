package model_test

import (
	"testing"

	"github.com/asztemborski/syncro/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	account := model.NewAccount("Test@example.com   ", "   Test_username")

	assert.NotEmpty(t, account.ID)
	assert.Equal(t, account.Email, "test@example.com")
	assert.Equal(t, account.Username, "test_username")
	assert.Equal(t, account.EmailVerified, false)
	assert.Equal(t, account.IsActive, true)
}

func TestPasswordSet(t *testing.T) {
	account := model.NewAccount("test@example.com", "test_username")
	err := account.HashPassword("secret_password")

	assert.NoError(t, err)
	assert.True(t, account.ComparePassword("secret_password"))
}
