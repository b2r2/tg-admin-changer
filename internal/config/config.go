package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"github.com/b2r2/tg-admin-changer/internal/models"
	"github.com/sirupsen/logrus"
)

var texts []string

func init() {
	files := []string{models.Greeting, models.Pricing, models.Repeated}
	texts = append(texts, files...)
}

type Config struct {
	bot      *bot
	logger   *logrus.Logger
	redis    *Redis
	database *Database
	mapping  map[string]*bytes.Buffer
}

var conf *Config

func Load() error {
	rd, err := NewRedis()
	if err != nil {
		return err
	}

	db, err := NewDB()
	if err != nil {
		return err
	}

	b, err := newBot()
	if err != nil {
		return err
	}

	conf = &Config{
		bot:      b,
		redis:    rd,
		database: db,
		logger:   logrus.New(),
		mapping:  make(map[string]*bytes.Buffer),
	}

	return conf.setTextFromFile()
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

func (c *Config) GetDBConnection() *sqlx.DB {
	return c.database.GetConnection()
}

func (c *Config) GetRedis() *redis.Client {
	return c.redis.GetClient()
}

func (c *Config) setTextFromFile() error {
	for _, filename := range texts {
		b, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("error set mapping text from file: %w", err)
		}

		c.mapping[filename] = bytes.NewBuffer(b)
	}

	return nil
}
