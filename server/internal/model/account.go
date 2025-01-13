package model

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID            uuid.UUID
	Email         string
	EmailVerified bool
	Username      string
	IsActive      bool
	Password      []byte
}

func NewAccount(email, username string) *Account {
	return &Account{
		ID:       uuid.New(),
		Email:    normalizeString(email),
		Username: normalizeString(username),
		IsActive: true,
	}
}

func (a *Account) HashPassword(value string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash a password: %w", err)
	}

	a.Password = hash
	return nil
}

func (a *Account) ComparePassword(value string) bool {
	return bcrypt.CompareHashAndPassword(a.Password, []byte(value)) == nil
}

func normalizeString(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}
