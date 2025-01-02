package model

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Email    string
	Username string
	Password password
}

func NewAccount(email, username string) *Account {
	return &Account{
		Email:    email,
		Username: username,
	}
}

func (a *Account) NormalizeEmail() {
	a.Email = normalizeString(a.Email)
}

func (a *Account) NormalizeUsername() {
	a.Username = normalizeString(a.Username)
}

type password struct {
	value *string
	hash  []byte
}

func (p *password) Set(value string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash a password: %w", err)
	}

	p.value = &value
	p.hash = hash
	return nil
}

func (p *password) Compare(value string) bool {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(value)) == nil
}

func normalizeString(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}
