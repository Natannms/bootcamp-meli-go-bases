package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// diferentes declarações de contextos

	// Create a context with cancel function
	ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	// c := context.WithValue(context.Background(), "key qualquer", "value")
	// valor := c.Value("key qualquer")

	// Listen for interrupt signal (e.g., Ctrl+C)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	go func() {
		fmt.Println("executando alguma coisa")
		<-sigCh
		fmt.Println("Received interrupt signal. Canceling operations...")
		cancel() // Cancel all operations
	}()

	// Start some goroutines representing ongoing operations
	go doSomething(ctx)
	go doSomething(ctx)

	// Wait for operations to complete or for the cancel signal
	<-ctx.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("All operations completed or canceled.")
}

func doSomething(ctx context.Context) {
	intChannel := make(chan int, 1)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation canceled.")
			return
		case <-intChannel:
			fmt.Println("alguma coisa")
		default:
			// Simulate some work
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}
