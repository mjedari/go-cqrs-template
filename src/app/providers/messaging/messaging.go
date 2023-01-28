package messaging

import (
	"context"
	"github.com/mjedari/go-cqrs-template/common"
)

type IMessageReceiver interface {
	Receive(ctx context.Context, messages []common.EventMessage) error
}

type IMessageSender interface {
	Send(ctx context.Context, messages []common.EventMessage) error
}

type IEventDispatcher interface {
	Dispatch(ctx context.Context, events []interface{}) error
	RegisterAtLeastOnce(name string, handler ISubscribable)
	RegisterAtMostOnce(name string, handler ISubscribable)
}

type IWorker interface {
	Loop() error
}

type ISubscribable = func(ctx context.Context, events []interface{}) error

type IEventMessage interface {
	GetPayload() []byte
	GetTopic() string
}
