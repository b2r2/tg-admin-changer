package config

import "github.com/b2r2/tg-admin-changer/pkg"

//const filename = ".env"
// TODO
const filename = "TOKEN"

type bot struct {
	token string
}

func newBot() (*bot, error) {
	token, err := pkg.GetEnv(filename)
	if err != nil {
		return nil, err
	}
	return &bot{
		token: token,
	}, nil
}
