// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package user_sql

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.selectAllUserStmt, err = db.PrepareContext(ctx, selectAllUser); err != nil {
		return nil, fmt.Errorf("error preparing query SelectAllUser: %w", err)
	}
	if q.selectOneUserStmt, err = db.PrepareContext(ctx, selectOneUser); err != nil {
		return nil, fmt.Errorf("error preparing query SelectOneUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.selectAllUserStmt != nil {
		if cerr := q.selectAllUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectAllUserStmt: %w", cerr)
		}
	}
	if q.selectOneUserStmt != nil {
		if cerr := q.selectOneUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectOneUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                DBTX
	tx                *sql.Tx
	selectAllUserStmt *sql.Stmt
	selectOneUserStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                tx,
		tx:                tx,
		selectAllUserStmt: q.selectAllUserStmt,
		selectOneUserStmt: q.selectOneUserStmt,
	}
}
