package postgresql

import (
	context "context"
	sql "database/sql"

	sqlx "github.com/jmoiron/sqlx"
)

type Ext interface {
	sqlx.ExtContext
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type txKeyType struct{}

var txKey = txKeyType{}

type Postgresql struct {
	DB *sqlx.DB
}

func NewPostgresql(db *sqlx.DB) *Postgresql {
	return &Postgresql{
		DB: db,
	}
}

func (p *Postgresql) GetExecutor(ctx context.Context) Ext {
	if tx, ok := ctx.Value(txKey).(*sqlx.Tx); ok {
		return tx
	}
	return p.DB
}
