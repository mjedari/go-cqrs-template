package wiring

import (
	"github.com/mjedari/go-cqrs-template/src/infra/providers/storage"
	"github.com/redis/go-redis/v9"
)

func (w *Wire) GetRedis() *redis.Client {
	return w.Redis
}

func (w *Wire) GetRedisInfra() *storage.RedisStorage {
	return storage.NewRedisStorage(w.GetRedis())
}
