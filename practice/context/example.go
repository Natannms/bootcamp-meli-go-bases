package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func example() {
	// Create a context with cancel function
	ctx, cancel := context.WithCancel(context.Background())

	// Listen for interrupt signal (e.g., Ctrl+C)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	go func() {
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
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation canceled.")
			return
		default:
			// Simulate some work
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}
