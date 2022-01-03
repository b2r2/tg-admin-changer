package repositories

import (
	"context"

	"github.com/b2r2/tg-admin-changer/internal/models"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetUser(ctx context.Context, username string) (*models.User, error)
	SetUser(ctx context.Context, first, username string) (*models.User, error)
	UpdateUser(ctx context.Context, id uint64) error
	Close() error
}

type repo struct {
	db    *sqlx.DB
	cache redis.UniversalClient
}

func NewRepository(db *sqlx.DB, r redis.UniversalClient) Repository {
	return &repo{
		db:    db,
		cache: r,
	}
}
