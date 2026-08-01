package application

import (
	"time"

	schedulerdomain "uptime-monitoring-platform/go-monolith/internal/scheduler/domain"
)

func PlanNextRun(schedule schedulerdomain.Schedule, now time.Time) time.Time {
	if schedule.Interval <= 0 {
		return now
	}

	return now.Add(schedule.Interval)
}
