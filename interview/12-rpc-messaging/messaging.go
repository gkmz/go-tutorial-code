// Package messaging 提供 gRPC、Kafka 和消息可靠性面试的最小模型。
package messaging

import (
	"context"
	"errors"
	"hash/fnv"
	"sync"
)

// UnaryCall 表示带 Deadline 的 RPC 调用模型。
func UnaryCall(ctx context.Context, work func(context.Context) error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return work(ctx)
}

// PartitionForKey 将相同 Key 稳定映射到同一 Partition。
func PartitionForKey(key string, partitions int) int {
	if partitions <= 0 {
		return -1
	}
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(key))
	return int(hash.Sum32() % uint32(partitions))
}

// Message 表示带业务 ID 的消息。
type Message struct {
	ID    string
	Key   string
	Value string
}

// Consumer 是一个支持重试和死信的内存消费者。
type Consumer struct {
	mu        sync.Mutex
	processed map[string]struct{}
	inFlight  map[string]*sync.Mutex
	dead      []Message
	maxRetry  int
}

// NewConsumer 创建消息消费者。
func NewConsumer(maxRetry int) *Consumer {
	if maxRetry < 0 {
		maxRetry = 0
	}
	return &Consumer{processed: make(map[string]struct{}), inFlight: make(map[string]*sync.Mutex), maxRetry: maxRetry}
}

// Consume 处理消息；重复消息不会再次执行副作用，失败消息超过次数后进入死信。
func (c *Consumer) Consume(message Message, handle func(Message) error) error {
	c.mu.Lock()
	if _, ok := c.processed[message.ID]; ok {
		c.mu.Unlock()
		return nil
	}
	keyLock, ok := c.inFlight[message.ID]
	if !ok {
		keyLock = &sync.Mutex{}
		c.inFlight[message.ID] = keyLock
	}
	c.mu.Unlock()

	// 同一消息 ID 串行处理，避免并发到达时重复执行副作用。
	keyLock.Lock()
	defer func() {
		keyLock.Unlock()
		c.mu.Lock()
		delete(c.inFlight, message.ID)
		c.mu.Unlock()
	}()
	c.mu.Lock()
	if _, ok := c.processed[message.ID]; ok {
		c.mu.Unlock()
		return nil
	}
	c.mu.Unlock()

	var err error
	for attempt := 0; attempt <= c.maxRetry; attempt++ {
		err = handle(message)
		if err == nil {
			c.mu.Lock()
			c.processed[message.ID] = struct{}{}
			c.mu.Unlock()
			return nil
		}
	}
	c.mu.Lock()
	c.dead = append(c.dead, message)
	c.mu.Unlock()
	return err
}

// DeadLetters 返回当前死信消息快照。
func (c *Consumer) DeadLetters() []Message {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]Message(nil), c.dead...)
}

// ValidateMessage 检查消息是否具备幂等键。
func ValidateMessage(message Message) error {
	if message.ID == "" {
		return errors.New("message id is required")
	}
	return nil
}
