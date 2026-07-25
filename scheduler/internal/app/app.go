package app
package app

import (
	"context"
	"fmt"
	"time"

	"uptime-monitoring-platform/scheduler/internal/config"
	"uptime-monitoring-platform/scheduler/internal/scheduler"
)

func Run(ctx context.Context) error {
	cfg := config.Load()
	engine := scheduler.New(cfg)

	fmt.Println("scheduler starting")
	return engine.Run(ctx)
}
