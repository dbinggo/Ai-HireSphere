package redisx

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func Open(cfg Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func MustOpen(cfg Redis) *redis.Client {
	client, err := Open(cfg)
	if err != nil {
		panic(err)
	}
	return client
}
