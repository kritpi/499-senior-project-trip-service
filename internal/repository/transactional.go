package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) Transactional(ctx context.Context, fn func(txCtx context.Context) error) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	// transactional-scoped context
	txCtx := context.WithValue(ctx, "tx", tx)

	if err := fn(txCtx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	committed = true
	return nil
}
