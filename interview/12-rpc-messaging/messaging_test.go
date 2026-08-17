package messaging

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
)

func TestUnaryCallHonorsDeadline(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := UnaryCall(ctx, func(context.Context) error { t.Fatal("work should not run"); return nil }); !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
}

func TestPartitionForKeyIsStable(t *testing.T) {
	first := PartitionForKey("BTC-USDT", 8)
	if first != PartitionForKey("BTC-USDT", 8) {
		t.Fatal("same key should map to same partition")
	}
	if PartitionForKey("BTC-USDT", 0) != -1 {
		t.Fatal("invalid partition count should return -1")
	}
}

func TestConsumerRetriesAndDeduplicates(t *testing.T) {
	consumer := NewConsumer(2)
	message := Message{ID: "event-1", Key: "key", Value: "value"}
	calls := 0
	err := consumer.Consume(message, func(Message) error {
		calls++
		if calls < 3 {
			return errors.New("temporary")
		}
		return nil
	})
	if err != nil || calls != 3 {
		t.Fatalf("err=%v calls=%d", err, calls)
	}
	if err := consumer.Consume(message, func(Message) error {
		t.Fatal("duplicate should not be handled")
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestConsumerDeadLetter(t *testing.T) {
	consumer := NewConsumer(1)
	message := Message{ID: "event-1"}
	if err := consumer.Consume(message, func(Message) error { return errors.New("permanent") }); err == nil {
		t.Fatal("expected permanent error")
	}
	if len(consumer.DeadLetters()) != 1 {
		t.Fatal("message should be in dead letters")
	}
}

func TestConsumerDeduplicatesConcurrentDelivery(t *testing.T) {
	consumer := NewConsumer(0)
	message := Message{ID: "event-concurrent"}
	var calls atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = consumer.Consume(message, func(Message) error {
				calls.Add(1)
				return nil
			})
		}()
	}
	wg.Wait()
	if calls.Load() != 1 {
		t.Fatalf("handle calls = %d, want 1", calls.Load())
	}
}
