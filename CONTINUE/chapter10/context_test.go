package chapter10

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	A := context.Background()
	B := context.WithValue(A, "b", "value B")
	C := context.WithValue(A, "c", "value C")
	D := context.WithValue(C, "d", "value D")

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)

	fmt.Println(A.Value("a")) // <nil> because A doenst have "value A"
	fmt.Println(D.Value("c")) // got from the parent
	fmt.Println(D.Value("d")) // got from herself
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	fmt.Println(runtime.NumGoroutine())

	for i := range destination {
		fmt.Println("Counter ", i)
		if i == 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine())
}

func CreateCounterSlow(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(2 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounterSlow(ctx)

	fmt.Println(runtime.NumGoroutine())

	for i := range destination {
		fmt.Println("Counter ", i)
		if i == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounterSlow(ctx)

	fmt.Println(runtime.NumGoroutine())

	for i := range destination {
		fmt.Println("Counter ", i)
		if i == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine())
}
