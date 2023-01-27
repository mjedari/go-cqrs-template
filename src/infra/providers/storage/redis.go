package storage

import (
	"context"
	"fmt"
	"github.com/mjedari/go-cqrs-template/src/api/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisStorage struct {
	*redis.Client
}

func NewRedis(conf config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Host, conf.Port),
		Username: conf.User,
		Password: conf.Pass,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisStorage(client *redis.Client) *RedisStorage {
	return &RedisStorage{Client: client}
}

func (r RedisStorage) Insert(ctx context.Context, key, value string) error {
	if err := r.Set(context.Background(), key, value, time.Minute).Err(); err != nil {
		return err
	}

	return nil
}
