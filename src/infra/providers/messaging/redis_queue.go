package messaging

import (
	"context"
	"errors"
	"fmt"
	"github.com/mjedari/go-cqrs-template/src/app/providers/messaging"
	"github.com/mjedari/go-cqrs-template/src/common"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var IdleWorkerError = errors.New("ideal workers error")

type RedisQueueReceiver struct {
	queueName       string
	receiver        messaging.IMessageReceiver
	redis           redis.Cmdable
	errorQueueDelay time.Duration
	popLimit        int
}

func NewRedisQueueReceiver(queueName string, receiver messaging.IMessageReceiver, redis redis.Cmdable) *RedisQueueReceiver {
	return &RedisQueueReceiver{queueName: queueName, receiver: receiver, redis: redis, popLimit: 1}
}

// register this loop or run it from root?
func (r *RedisQueueReceiver) Loop() error {
	ctx := context.Background()
	popCmd := r.redis.ZPopMin(ctx, r.queueName, int64(r.popLimit))
	if err := popCmd.Err(); err != nil {
		return err
	}

	items := popCmd.Val()
	if len(items) == 0 {
		return IdleWorkerError
	}

	messages := make([]common.EventMessage, 0)
	for idx, item := range items {
		if item.Score > float64(time.Now().Unix()) {
			if err := r.redis.ZAdd(ctx, r.queueName, items[idx:]...).Err(); err != nil {
				//redisLogger.Error(err)
			}

			if idx == 0 {
				return IdleWorkerError
			} else {
				break
			}
		}

		payload := []byte(item.Member.(string))
		messages = append(messages, common.EventMessage{
			Payload: payload,
			Topic:   r.queueName,
		})
	}

	locker := sync.WaitGroup{}
	var errors error
	runFunc := func(message common.EventMessage) {
		defer locker.Done()
		if err := r.receiver.Receive(context.Background(), []common.EventMessage{message}); err != nil {
			// if error happens add it to queue again

		}
	}

	for _, message := range messages {
		locker.Add(1)
		go runFunc(message)
	}

	locker.Wait()
	return errors
}

type RedisQueueSender struct {
	redis redis.Cmdable
}

func NewRedisQueueSender(redis redis.Cmdable) *RedisQueueSender {
	return &RedisQueueSender{redis: redis}
}

func (r *RedisQueueSender) Send(ctx context.Context, messages []common.EventMessage) error {
	fmt.Println("Send called: ", string(messages[0].Payload))
	// todo: How to publish payload into redis

	members := make(map[string][]redis.Z)
	for _, message := range messages {
		if _, ok := members[message.Topic]; !ok {
			members[message.Topic] = make([]redis.Z, 0)
		}
		members[message.Topic] = append(members[message.Topic], redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: message.Payload,
		})
	}
	for topic, topicMembers := range members {
		if err := r.redis.ZAdd(ctx, topic, topicMembers...).Err(); err != nil {
			return err
		}
	}

	return nil
}
