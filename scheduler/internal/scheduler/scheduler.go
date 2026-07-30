package main
package scheduler

import (
	"context"
	"time"

	"uptime-monitoring-platform/scheduler/internal/config"
)

type Scheduler struct {
	cfg config.Config
}

func New(cfg config.Config) *Scheduler {
	return &Scheduler{cfg: cfg}
}

func (s *Scheduler) Run(ctx context.Context) error {
	interval, err := time.ParseDuration(s.cfg.Interval)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			// place scheduled jobs here
		}
	}
}
