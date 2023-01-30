package storage

import (
	"context"
	"github.com/mjedari/go-cqrs-template/domain/contract"
)

type IStorage interface {
	Insert(ctx context.Context, key string, value []byte) error
	Select(ctx context.Context, key string) ([]byte, error)
	SelectAll(ctx context.Context, match string) ([][]byte, error)
	Update(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error

	GetNextId(ctx context.Context, entity contract.IEntity) (int64, error)
}
