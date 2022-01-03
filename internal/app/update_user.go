package app

import (
	"context"
)

func (b *bot) UpdateUser(ctx context.Context, id uint64) error {
	return b.repo.UpdateUser(ctx, id)
}
