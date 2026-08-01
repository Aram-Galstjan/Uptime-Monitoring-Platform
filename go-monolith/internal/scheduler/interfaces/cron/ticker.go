package cron

import "context"

type Scheduler interface {
	Tick(ctx context.Context) error
}
