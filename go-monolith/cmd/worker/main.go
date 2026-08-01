package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"uptime-monitoring-platform/go-monolith/internal/checker/interfaces/worker"
)

func main() {
	if err := run(); err != nil {
		log.Printf("worker failed: %v", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := worker.RunWorker(ctx); err != nil && err != context.Canceled {
		return fmt.Errorf("run worker: %w", err)
	}

	return nil
}
