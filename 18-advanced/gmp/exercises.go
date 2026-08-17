package main

import (
	"context"
	"errors"
	"sync"
)

func runLimited(ctx context.Context, tasks, limit int, work func(int)) error {
	if ctx == nil {
		return errors.New("context must not be nil")
	}
	if tasks < 0 {
		return errors.New("tasks must not be negative")
	}
	if limit < 1 {
		return errors.New("limit must be greater than zero")
	}
	if work == nil {
		return errors.New("work must not be nil")
	}

	jobs := make(chan int)
	var workers sync.WaitGroup
	for range min(tasks, limit) {
		workers.Go(func() {
			for index := range jobs {
				select {
				case <-ctx.Done():
					return
				default:
					work(index)
				}
			}
		})
	}

	for index := 0; index < tasks; index++ {
		select {
		case jobs <- index:
		case <-ctx.Done():
			close(jobs)
			workers.Wait()
			return ctx.Err()
		}
	}
	close(jobs)
	workers.Wait()
	return ctx.Err()
}
