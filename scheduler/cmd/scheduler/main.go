package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"uptime-monitoring-platform/scheduler/internal/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx); err != nil {
		log.Println("scheduler exited with error:", err)
		os.Exit(1)
	}
}
