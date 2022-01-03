package app

import (
	"context"
	"database/sql"
	"errors"

	"github.com/b2r2/tg-admin-changer/internal/models"
)

func (b *bot) GetUser(ctx context.Context, username string) (u *models.User, err error) {
	if u, err = b.repo.GetUser(ctx, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return u, err
}
