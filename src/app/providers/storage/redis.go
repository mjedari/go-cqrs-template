package storage

import "context"

type IRedisStorage interface {
	Insert(ctx context.Context, key, value string) error
}
