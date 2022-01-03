package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisConnection struct {
	Host       string
	Port       string
	Password   string
	DB         int
	MasterName string
	Sentinel   []string
}

type Redis struct {
	connectionConfig RedisConnection
	client           *redis.Client
}

func (r *Redis) GetClient() *redis.Client {
	return r.client
}

func NewRedis() (*Redis, error) {
	cfg := RedisConnection{
		// TODO rename localhost to redis
		Host:       getEnv("REDIS_HOST", "localhost"),
		Port:       getEnv("REDIS_PORT", "6379"),
		Password:   getEnv("REDIS_PASSWORD", ""),
		DB:         getEnvAsInt("REDIS_DB", 0),
		MasterName: getEnv("REDIS_MASTER_NAME", "mymaster"),
		Sentinel:   getEnvAsSlice("REDIS_SENTINEL", make([]string, 0), ","),
	}

	var client *redis.Client
	if len(cfg.Sentinel) > 0 {
		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    cfg.MasterName,
			SentinelAddrs: cfg.Sentinel,
		})
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:       fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Password:   cfg.Password,
			DB:         cfg.DB,
			MaxRetries: 3,
		})
	}
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	r := Redis{
		connectionConfig: cfg,
		client:           client,
	}

	return &r, nil
}
