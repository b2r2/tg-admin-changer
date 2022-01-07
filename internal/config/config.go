package config

import (
	"bytes"

	"github.com/b2r2/tg-admin-changer/internal/models"
	"github.com/sirupsen/logrus"
)

type Config struct {
	bot     *bot
	logger  *logrus.Logger
	mapping map[string]*bytes.Buffer
}

var conf *Config

func Load() error {
	b, err := newBot()
	if err != nil {
		return err
	}

	conf = &Config{
		bot:     b,
		logger:  logrus.New(),
		mapping: models.Texts,
	}

	return nil
}

func Get() *Config {
	return conf
}

func (c *Config) GetToken() string {
	return conf.bot.token
}

func (c *Config) GetLogger() *logrus.Logger {
	return conf.logger
}

func (c *Config) GetMapping() map[string]*bytes.Buffer {
	return conf.mapping
}
