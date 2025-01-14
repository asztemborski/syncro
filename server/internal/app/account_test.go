package app_test

import (
	"context"
	"testing"

	"github.com/asztemborski/syncro/internal/app"
	"github.com/asztemborski/syncro/internal/model"
	"github.com/stretchr/testify/assert"
)

type InMemoryAccountStore struct {
	accounts []model.Account
}

func NewInMemoryAccountStore() *InMemoryAccountStore {
	return &InMemoryAccountStore{
		accounts: make([]model.Account, 0),
	}
}

func (s *InMemoryAccountStore) Save(_ context.Context, account *model.Account) error {
	s.accounts = append(s.accounts, *account)
	return nil
}

func (s *InMemoryAccountStore) IsUnique(_ context.Context, account *model.Account) (bool, bool) {
	for _, acc := range s.accounts {
		switch {
		case acc.Email == account.Email:
			return false, true
		case acc.Username == account.Username:
			return true, false
		}
	}

	return true, true
}

func (s *InMemoryAccountStore) Add(_ context.Context, account *model.Account) {
	s.accounts = append(s.accounts, *account)
}

func (s *InMemoryAccountStore) Clear() {
	s.accounts = []model.Account{}
}

func TestAccountService_CreateAccount(t *testing.T) {
	store := NewInMemoryAccountStore()
	service := app.NewAccountService(store)
	ctx := context.Background()

	t.Run("should save new account if email and username are unique", func(t *testing.T) {
		store.Clear()
		store.Add(ctx, model.NewAccount("test@example.com", "testuser"))
		payload := app.CreateAccountPayload{
			Email:    "test@unique.com",
			Username: "testunique",
			Password: "password123",
		}

		err := service.CreateAccount(ctx, payload)
		assert.NoError(t, err)
	})

	t.Run("should return error if email is not unique", func(t *testing.T) {
		store.Clear()
		store.Add(ctx, model.NewAccount("test@example.com", "testuser"))
		payload := app.CreateAccountPayload{
			Email:    "test@example.com",
			Username: "testunique",
			Password: "password123",
		}

		err := service.CreateAccount(ctx, payload)
		assert.Equal(t, err, app.ErrEmailInUse)
	})

	t.Run("should return error if username is not unique", func(t *testing.T) {
		store.Clear()
		store.Add(ctx, model.NewAccount("test@example.com", "testuser"))
		payload := app.CreateAccountPayload{
			Email:    "test@unique.com",
			Username: "testuser",
			Password: "password123",
		}

		err := service.CreateAccount(ctx, payload)
		assert.Equal(t, err, app.ErrUsernameInUse)
	})
}
