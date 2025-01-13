package store

import (
	"context"
	"errors"

	"github.com/asztemborski/syncro/internal/model"
	"github.com/asztemborski/syncro/internal/store/postgres/public/table"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
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

func (s *SqlAccountStore) IsUnique(ctx context.Context, account *model.Account) (bool, bool) {
	stmt := pg.SELECT(
		pg.CASE().
			WHEN(pg.EXISTS(
				pg.SELECT(table.Account.ID).
					FROM(table.Account).
					WHERE(table.Account.Email.EQ(pg.String(account.Email))),
			)).
			THEN(pg.Bool(false)).
			ELSE(pg.Bool(true)).
			AS("email_unique"),
		pg.CASE().
			WHEN(pg.EXISTS(
				pg.SELECT(table.Account.ID).
					FROM(table.Account).
					WHERE(table.Account.Username.EQ(pg.String(account.Username))),
			)).
			THEN(pg.Bool(false)).
			ELSE(pg.Bool(true)).
			AS("username_unique"),
	).FROM(table.Account).LIMIT(1)

	var result struct {
		EmailUnique    bool `db:"email_unique"`
		UsernameUnique bool `db:"username_unique"`
	}

	if err := stmt.QueryContext(ctx, s.db, &result); err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return true, true
		}
		return false, false
	}

	return result.EmailUnique, result.UsernameUnique
}
