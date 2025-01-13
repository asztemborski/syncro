package store

import (
	"context"

	"github.com/asztemborski/syncro/internal/model"
	"github.com/jmoiron/sqlx"
)

type AccountStore interface {
	Save(context.Context, *model.Account) error
	IsUnique(context.Context, *model.Account) (bool, bool)
}

type Store interface {
	Account() AccountStore
}

type SqlStore struct {
	account *SqlAccountStore
}

func NewSqlStore(db *sqlx.DB) *SqlStore {
	return &SqlStore{
		account: &SqlAccountStore{db: db},
	}
}

func (s *SqlStore) Account() AccountStore {
	return s.account
}
