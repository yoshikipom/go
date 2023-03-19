package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

func main() {
	// goroutine()
	// anonymousGoroutine()
	// waitTimer()
	// forWithChannel()
	// contextTest()
	timerTest()
}

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func goroutine() {
	fmt.Println("start sub()")
	go sub()
	time.Sleep(2 * time.Second)
}

func anonymousGoroutine() {
	fmt.Println("start sub()")
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
}

func waitTimer() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		done <- true
	}()
	<-done
	fmt.Println("all tasks are finished")
}

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 1000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l+1; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func forWithChannel() {
	pn := primeNumber()
	for n := range pn {
		fmt.Println(n)
	}
}

func contextTest() {
	fmt.Println("start sub()")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("all tasks are finished")

	fmt.Println(<-time.After(3 * time.Second))
	fmt.Println("after timer")
}

func timerTest() {
	fmt.Println(<-time.After(3 * time.Second))
	fmt.Println("after timer")
}
