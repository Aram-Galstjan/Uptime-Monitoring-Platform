package application

import (
	"context"

	"uptime-monitoring-platform/checker/internal/domain"
)

type Runner interface {
	Run(ctx context.Context, target string) (domain.CheckResult, error)
}