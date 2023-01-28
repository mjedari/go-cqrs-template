package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mjedari/go-cqrs-template/app/providers/messaging"
	"github.com/mjedari/go-cqrs-template/common"
	"github.com/redis/go-redis/v9"
	"reflect"
)

type EventBus struct {
	dispatcher *EventDispatcher
	publisher  *JsonEventPublisher
	receiver   *RedisQueueReceiver
	registry   *EventRegistry
}

func (e *EventBus) Loop() error {
	return e.receiver.Loop()
}

func NewEventBus(topic string, redisClient redis.Cmdable) *EventBus {
	dispatcher := NewEventDispatcher()
	sender := NewRedisQueueSender(redisClient)
	registry := NewEventRegistry()
	publisher := NewJsonEventPublisher(sender, registry, topic)
	messageReceiver := NewJsonMessageReceiver(dispatcher, registry)
	receiver := NewRedisQueueReceiver(topic, messageReceiver, redisClient)

	return &EventBus{
		dispatcher: dispatcher,
		publisher:  publisher,
		receiver:   receiver,
		registry:   registry,
	}
}

type EventDispatcher struct {
	atMostOnceSubscribers map[string]messaging.ISubscribable
}

func (e EventDispatcher) Dispatch(ctx context.Context, events []interface{}) error {
	e.dispatchAtMostOnce(ctx, events)
	return nil
}

func (e *EventDispatcher) dispatchAtMostOnce(ctx context.Context, events []interface{}) {
	for name, subscriber := range e.atMostOnceSubscribers {
		if err := subscriber(ctx, events); err != nil {
			// err
			println(name, err)
		}
	}
}

func (e EventDispatcher) RegisterAtLeastOnce(name string, handler messaging.ISubscribable) {
	//TODO implement me
	panic("implement me")
}

func (e EventDispatcher) RegisterAtMostOnce(name string, handler messaging.ISubscribable) {
	//TODO implement me
	panic("implement me")
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		atMostOnceSubscribers: make(map[string]messaging.ISubscribable),
	}
}

type JsonEventPublisher struct {
	messageSender messaging.IMessageSender
	eventRegistry *EventRegistry
	topicName     string
}

func NewJsonEventPublisher(messageSender messaging.IMessageSender, eventRegistry *EventRegistry, topicName string) *JsonEventPublisher {
	return &JsonEventPublisher{messageSender: messageSender, eventRegistry: eventRegistry, topicName: topicName}
}

type JsonMessageReceiver struct {
	eventDispatcher messaging.IEventDispatcher
	eventRegistry   *EventRegistry
}

func (j JsonMessageReceiver) Receive(ctx context.Context, messages []common.EventMessage) error {
	events := make([]interface{}, 0)
	for _, message := range messages {
		typedEvent := struct {
			Name string
		}{}
		if err := json.Unmarshal(message.Payload, &typedEvent); err != nil {
			return err
		}

		eventType, err := j.eventRegistry.GetByName(typedEvent.Name)
		if err != nil {
			return err
		}

		eventWrapper := reflect.New(reflect.StructOf([]reflect.StructField{
			{Name: "Event", Type: eventType},
		}))

		if err := json.Unmarshal(message.Payload, eventWrapper.Interface()); err != nil {
			return err
		}

		event := eventWrapper.Elem().FieldByName("Event").Interface()
		events = append(events, event)
	}

	return j.eventDispatcher.Dispatch(ctx, events)
}

func NewJsonMessageReceiver(eventDispatcher messaging.IEventDispatcher, eventRegistry *EventRegistry) *JsonMessageReceiver {
	return &JsonMessageReceiver{eventDispatcher: eventDispatcher, eventRegistry: eventRegistry}
}

func (e *JsonEventPublisher) Receive(ctx context.Context, messages []common.EventMessage) error {
	fmt.Println("receiver", messages)
	return nil
}

func (bus EventBus) Dispatch(ctx context.Context, event interface{}) error {

	fmt.Println("Event bus")
	return nil
}

func (bus EventBus) RegisterListener(s string, f func(ctx context.Context, events []interface{}) error) {
	bus.dispatcher.atMostOnceSubscribers[s] = f
}

func (e *JsonEventPublisher) Publish(ctx context.Context, events []interface{}) error {
	messages := make([]common.EventMessage, 0)

	for _, event := range events {
		name, err := e.eventRegistry.GetNameByType(reflect.TypeOf(event))
		payload, err := json.Marshal(struct {
			Name  string
			Event interface{}
		}{
			Name:  name,
			Event: event,
		})
		if err != nil {
			return err
		}

		messages = append(messages, common.EventMessage{
			Payload: payload,
			Topic:   e.topicName,
		})
	}

	return e.messageSender.Send(ctx, messages)
}

type EventRegistry struct {
	typesByName map[string]reflect.Type
	namesByType map[reflect.Type]string
}

func NewEventRegistry() *EventRegistry {
	return &EventRegistry{
		typesByName: make(map[string]reflect.Type),
		namesByType: make(map[reflect.Type]string),
	}
}

func (r *EventRegistry) Register(event interface{}) *EventRegistry {
	name := fmt.Sprintf("%s.%s", reflect.TypeOf(event).PkgPath(), reflect.TypeOf(event).Name())
	return r.RegisterByName(name, event)
}

func (r *EventRegistry) RegisterByName(name string, event interface{}) *EventRegistry {
	r.typesByName[name] = reflect.TypeOf(event)
	r.namesByType[reflect.TypeOf(event)] = name

	return r
}

func (r *EventRegistry) GetByName(name string) (reflect.Type, error) {
	ev, ok := r.typesByName[name]
	if !ok {
		return nil, nil
	}

	return ev, nil
}

func (r *EventRegistry) GetNameByType(typ reflect.Type) (string, error) {
	ev, ok := r.namesByType[typ]
	if !ok {
		return "", nil
	}

	return ev, nil
}

func (e *EventBus) Publish(ctx context.Context, events []interface{}) error {
	return e.publisher.Publish(ctx, events)
}

func (e *EventBus) RegisterEvent(event interface{}) *EventBus {
	e.registry.Register(event)
	return e
}
