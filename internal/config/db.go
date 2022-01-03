package config

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConnectionConfig struct {
	Driver string
	Host   string
	Port   string
	User   string
	Pass   string
	Name   string
}

type Database struct {
	connectionConfig *ConnectionConfig
	connection       *sqlx.DB
}

func (cc *ConnectionConfig) GetDSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		cc.Driver,
		cc.User,
		cc.Pass,
		cc.Host,
		cc.Port,
		cc.Name,
	)
}

func (d *Database) GetConnection() *sqlx.DB {
	return d.connection
}

func NewDB() (*Database, error) {
	cfg := ConnectionConfig{
		Driver: getEnv("DB_DRIVER", "postgres"),
		// TODO rename localhost to postgres
		Host: getEnv("DB_HOST", "localhost"),
		Port: getEnv("DB_PORT", "5432"),
		User: getEnv("DB_USERNAME", "postgres"),
		Pass: getEnv("DB_PASSWORD", "postgres"),
		Name: getEnv("DB_DATABASE", "postgres"),
	}

	conn, err := sqlx.Connect(cfg.Driver, cfg.GetDSN())
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	d := Database{
		connectionConfig: &cfg,
		connection:       conn,
	}

	return &d, nil
}
