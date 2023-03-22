package main_test

import (
	"context"
	"fmt"
	"testing"
)

type ctxUserID struct{}
type ctxAuthToken struct{}
type ctxTraceID struct{}

func valueGenerator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num:
			}
		}
		close(out)
		userID, authToken, traceID := ctx.Value(ctxUserID{}).(int), ctx.Value(ctxAuthToken{}).(string), ctx.Value(ctxTraceID{}).(int)
		fmt.Println("log: ", userID, authToken, traceID)
	}()
	return out
}

func TestContextValue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, ctxUserID{}, 2)
	ctx = context.WithValue(ctx, ctxAuthToken{}, "xxxxx")
	ctx = context.WithValue(ctx, ctxTraceID{}, 3)
	gen := valueGenerator(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	cancel()

	wg.Wait()
}
