package app

import (
	"context"

	"github.com/asztemborski/syncro/internal/model"
	"github.com/asztemborski/syncro/internal/store"
)

var (
	ErrUsernameInUse = model.NewAppErr("account.username_used", "username already in use")
	ErrEmailInUse    = model.NewAppErr("account.email_used", "email already in use")
)

type AccountService struct {
	accountStore store.AccountStore
}

func NewAccountService(accountStore store.AccountStore) *AccountService {
	return &AccountService{accountStore: accountStore}
}

type CreateAccountPayload struct {
	Email    string
	Username string
	Password string
}

func (as *AccountService) CreateAccount(ctx context.Context, payload CreateAccountPayload) error {
	if err := as.validateUniqueConstraints(ctx, payload.Email, payload.Username); err != nil {
		return err
	}

	account := model.NewAccount(payload.Email, payload.Username)
	if err := account.HashPassword(payload.Password); err != nil {
		return err
	}

	return as.accountStore.Save(ctx, account)
}

func (as *AccountService) validateUniqueConstraints(ctx context.Context, email, username string) error {
	isEmailUnique, isUsernameUnique := as.accountStore.IsUnique(ctx, &model.Account{
		Email:    email,
		Username: username,
	})

	switch {
	case !isUsernameUnique:
		return ErrUsernameInUse
	case !isEmailUnique:
		return ErrEmailInUse
	}

	return nil
}
