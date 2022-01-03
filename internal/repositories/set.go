package repositories

import (
	"context"

	"github.com/b2r2/tg-admin-changer/internal/models"
)

func (r *repo) SetUser(ctx context.Context, first, username string) (*models.User, error) {
	var u models.User
	if err := r.db.GetContext(ctx, &u, models.SetUser, first, username); err != nil {
		return nil, err
	}

	go func() {
		_ = r.cache.Del(context.Background(), models.KeyCache)
	}()

	return &u, nil
}
