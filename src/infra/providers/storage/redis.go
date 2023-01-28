package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/mjedari/go-cqrs-template/domain/contract"
	"github.com/mjedari/go-cqrs-template/web/config"
	"github.com/redis/go-redis/v9"
	"reflect"
	"strings"
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

func (r RedisStorage) GetNextId(ctx context.Context, entity contract.IEntity) (int64, error) {
	entityType := reflect.TypeOf(entity)
	name := entityType.Elem().Name()

	key := fmt.Sprintf("next_%v_id", strings.ToLower(name))
	nextID, err := r.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return nextID, nil
}

func (r RedisStorage) Insert(ctx context.Context, key string, value []byte) error {
	if err := r.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (r RedisStorage) Select(ctx context.Context, key string) ([]byte, error) {
	fetched, err := r.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, errors.New("resource not found")
	}

	if err != nil {
		return nil, err
	}

	return fetched, nil
}

func (r RedisStorage) Update(ctx context.Context, key string, value []byte) error {
	if err := r.Delete(ctx, key); err != nil {
		return err
	}

	if err := r.Insert(ctx, key, value); err != nil {
		return err
	}
	return nil
}

func (r RedisStorage) Delete(ctx context.Context, key string) error {
	return r.Del(ctx, key).Err()
}
