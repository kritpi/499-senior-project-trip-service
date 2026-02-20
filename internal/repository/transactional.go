package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type txKey struct{}

func (r *Repository) Transactional(ctx context.Context, fn func(txCtx context.Context) error) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	// Defer a function that handles both panics and rollbacks
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p) // re-throw panic after rollback
		}
	}()

	txCtx := context.WithValue(ctx, txKey{}, tx)

	if err := fn(txCtx); err != nil {
		// If the user's function returns an error, we roll back
		_ = tx.Rollback(ctx)
		return err
	}

	// If we reached here, try to commit
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
