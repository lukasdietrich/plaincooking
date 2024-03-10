package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lukasdietrich/plaincooking/internal/database/models"
)

var (
	ErrNestedTx = errors.New("cannot begin transaction inside a running transaction")
)

type Querier interface {
	models.Querier
	Begin(context.Context) (Querier, *sql.Tx, error)
}

type querier struct {
	models.Querier
	db models.DBTX
}

func NewQuerier(db models.DBTX) Querier {
	return querier{
		Querier: models.New(db),
		db:      db,
	}
}

func (q querier) Begin(ctx context.Context) (Querier, *sql.Tx, error) {
	txer, ok := q.db.(interface {
		BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	})
	if !ok {
		return nil, nil, ErrNestedTx
	}

	tx, err := txer.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return NewQuerier(tx), tx, nil

}
