package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"uptime-monitoring-platform/scheduler/internal/app"
)

func main() {
	if err := run(); err != nil {
		log.Printf("scheduler failed: %v", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx); err != nil && err != context.Canceled {
		return fmt.Errorf("run app: %w", err)
	}

	return nil
}
