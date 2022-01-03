package repositories

import (
	"context"
	"database/sql"

	"github.com/b2r2/tg-admin-changer/internal/models"
)

func (r *repo) UpdateUser(ctx context.Context, id uint64) error {
	go func() {
		_ = r.cache.Del(context.Background(), models.KeyCache)
	}()

	res, err := r.db.ExecContext(ctx, models.UpdateUser, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
