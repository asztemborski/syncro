//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Account = newAccountTable("public", "account", "")

type accountTable struct {
	postgres.Table

	// Columns
	ID            postgres.ColumnString
	Email         postgres.ColumnString
	EmailVerified postgres.ColumnBool
	Username      postgres.ColumnString
	IsActive      postgres.ColumnBool
	Password      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AccountTable struct {
	accountTable

	EXCLUDED accountTable
}

// AS creates new AccountTable with assigned alias
func (a AccountTable) AS(alias string) *AccountTable {
	return newAccountTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AccountTable with assigned schema name
func (a AccountTable) FromSchema(schemaName string) *AccountTable {
	return newAccountTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AccountTable with assigned table prefix
func (a AccountTable) WithPrefix(prefix string) *AccountTable {
	return newAccountTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AccountTable with assigned table suffix
func (a AccountTable) WithSuffix(suffix string) *AccountTable {
	return newAccountTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAccountTable(schemaName, tableName, alias string) *AccountTable {
	return &AccountTable{
		accountTable: newAccountTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newAccountTableImpl("", "excluded", ""),
	}
}

func newAccountTableImpl(schemaName, tableName, alias string) accountTable {
	var (
		IDColumn            = postgres.StringColumn("id")
		EmailColumn         = postgres.StringColumn("email")
		EmailVerifiedColumn = postgres.BoolColumn("email_verified")
		UsernameColumn      = postgres.StringColumn("username")
		IsActiveColumn      = postgres.BoolColumn("is_active")
		PasswordColumn      = postgres.StringColumn("password")
		allColumns          = postgres.ColumnList{IDColumn, EmailColumn, EmailVerifiedColumn, UsernameColumn, IsActiveColumn, PasswordColumn}
		mutableColumns      = postgres.ColumnList{EmailColumn, EmailVerifiedColumn, UsernameColumn, IsActiveColumn, PasswordColumn}
	)

	return accountTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		Email:         EmailColumn,
		EmailVerified: EmailVerifiedColumn,
		Username:      UsernameColumn,
		IsActive:      IsActiveColumn,
		Password:      PasswordColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
