package config

import (
	"github.com/b2r2/tg-admin-changer/internal/models"
	"github.com/b2r2/tg-admin-changer/pkg"
)

type bot struct {
	token string
}

func newBot() (*bot, error) {
	token, err := pkg.GetEnv(models.Filename)
	if err != nil {
		return nil, err
	}
	return &bot{
		token: token,
	}, nil
}
