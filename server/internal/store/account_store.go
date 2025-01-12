package store

import (
	"context"

	"github.com/asztemborski/syncro/internal/model"
	"github.com/asztemborski/syncro/internal/store/postgres/public/table"
	"github.com/jmoiron/sqlx"
)

type SqlAccountStore struct {
	db *sqlx.DB
}

func (s *SqlAccountStore) Save(ctx context.Context, account *model.Account) error {
	stmt := table.Account.INSERT(table.Account.AllColumns).MODEL(account)
	_, err := stmt.ExecContext(ctx, s.db)
	return err
}
