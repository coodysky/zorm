package orm

import (
	"context"
	"database/sql"
	"github.com/envzo/zorm/db"
)

type Ztx struct {
	tx *sql.Tx
}

func TxBegin() (*Ztx, error) {
	tx, err := db.DB().Begin()
	if err != nil {
		return nil, err
	}

	txEntity := &Ztx{}
	txEntity.tx = tx

	return txEntity, nil
}

func (ztx *Ztx) Stmt(stmt *sql.Stmt) *sql.Stmt {
	return ztx.tx.Stmt(stmt)
}

func (ztx *Ztx) StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt {
	return ztx.tx.StmtContext(ctx, stmt)
}

func (ztx *Ztx) Prepare(query string) (*sql.Stmt, error) {
	return ztx.tx.Prepare(query)
}

func (ztx *Ztx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return ztx.tx.PrepareContext(ctx, query)
}

func (ztx *Ztx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return ztx.tx.Query(query, args...)
}

func (ztx *Ztx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return ztx.tx.QueryContext(ctx, query, args...)
}

func (ztx *Ztx) QueryRow(query string, args ...interface{}) *sql.Row {
	return ztx.tx.QueryRow(query, args...)
}

func (ztx *Ztx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return ztx.tx.QueryRowContext(ctx, query, args...)
}

func (ztx *Ztx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return ztx.tx.Exec(query, args...)
}

func (ztx *Ztx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return ztx.tx.ExecContext(ctx, query, args...)
}

func (ztx *Ztx) Commit() error {
	return ztx.tx.Commit()
}

func (ztx *Ztx) Rollback() error {
	return ztx.tx.Rollback()
}
