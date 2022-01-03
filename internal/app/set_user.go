package app

import (
	"context"
)

func (b *bot) SetUser(ctx context.Context, fn, un string) error {
	_, err := b.repo.SetUser(ctx, fn, un)
	return err
}
