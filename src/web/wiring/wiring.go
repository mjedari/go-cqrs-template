package wiring

import (
	"context"
	"github.com/mjedari/go-cqrs-template/domain/order"
	"github.com/mjedari/go-cqrs-template/infra/providers/messaging"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Wiring *Wire

type Wire struct {
	Redis    *redis.Client
	eventBus *messaging.EventBus
}

func NewWire(redis *redis.Client) *Wire {
	w := &Wire{Redis: redis}

	w.init()
	return w
}

func (w *Wire) init() {
	w.initEventBus()
	go w.registerWorker()
}

func (w *Wire) initEventBus() {
	w.eventBus = w.GetEventBus()
	w.eventBus.RegisterEvent(order.TestEvent{})
	w.eventBus.RegisterEvent(order.OrderEvent{})
	w.eventBus.RegisterEvent(order.InstantOrderEvent{})
	w.eventBus.RegisterEvent(order.FailTransactionEvent{})

	w.eventBus.RegisterListener("order-event", func(ctx context.Context, events []interface{}) error {
		if err := w.GetOrderEventHandler().HandleEvents(ctx, events); err != nil {
			return err
		}
		return nil
	})
}

func (w *Wire) registerWorker() {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-interrupt:
			logrus.Info("Terminating %s.")
			return
		case <-time.After(time.Second):
			if err := w.eventBus.Loop(); err != nil {
				//
				logrus.Debug("no queue wating")
			}
		}

	}
}
