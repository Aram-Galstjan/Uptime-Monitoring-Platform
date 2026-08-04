package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"uptime-monitoring-platform/go-monolith/internal/scheduler/interfaces/cron"
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

	if err := cron.RunScheduler(ctx); err != nil && err != context.Canceled {
		return fmt.Errorf("run scheduler: %w", err)
	}

	return nil
}
