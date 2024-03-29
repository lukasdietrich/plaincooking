package service

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/lukasdietrich/plaincooking/internal/database/models"
)

type txContextKey struct{}

type TransactionService struct {
	db *sql.DB
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{
		db: db,
	}
}

func (s *TransactionService) Transactional(ctx context.Context, fn func(context.Context) error) error {
	slog.Debug("beginning database transaction")

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, txContextKey{}, tx)

	if err := fn(ctx); err != nil {
		slog.Debug("rolling back database transaction")

		tx.Rollback() // nolint:errcheck
		return err
	}

	slog.Debug("committing database transaction")
	return tx.Commit()
}

func (s *TransactionService) Querier(ctx context.Context) models.Querier {
	tx, ok := ctx.Value(txContextKey{}).(*sql.Tx)
	if ok && tx != nil {
		return models.New(tx)
	}

	return models.New(s.db)
}
