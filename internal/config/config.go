package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/b2r2/tg-admin-changer/internal/models"
	"github.com/sirupsen/logrus"
)

var texts []string

func init() {
	files := []string{models.Greeting, models.Pricing, models.Contacts}
	texts = append(texts, files...)
}

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
		mapping: make(map[string]*bytes.Buffer),
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
