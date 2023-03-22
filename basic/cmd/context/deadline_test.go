package main_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var deadlineWg sync.WaitGroup

func deadlineGenerator(ctx context.Context, num int) <-chan int {
	out := make(chan int)

	go func() {
		defer deadlineWg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			}
		}
		close(out)
		fmt.Println("deadlineGenerator closed")
	}()
	return out
}

func TestDeadlineContext(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	gen := deadlineGenerator(ctx, 1)

	deadlineWg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		select {
		case result, ok := <-gen:
			if ok {
				fmt.Println(result)
			} else {
				fmt.Println("timeout")
				break LOOP
			}
		}
	}

	cancel()
}
