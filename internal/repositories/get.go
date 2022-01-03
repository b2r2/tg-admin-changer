package repositories

import (
	"context"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/b2r2/tg-admin-changer/internal/models"
)

func (r *repo) GetUser(ctx context.Context, username string) (*models.User, error) {
	var (
		us   models.Users
		user models.User
	)

	if data, err := r.cache.Get(ctx, models.KeyCache).Bytes(); err == nil {
		if err = jsoniter.Unmarshal(data, &us); err == nil {
			for _, u := range us {
				if u.Username == username {
					fmt.Println("redis done")
					return &u, nil
				}
			}
		}
	}

	if err := r.db.GetContext(ctx, &user, models.GetUser, username); err != nil {
		return nil, err
	}

	go func() {
		_ = r.cache.Set(context.Background(), models.KeyCache, user, time.Hour*720)
	}()

	return &user, nil
}
