package application

import "uptime-monitoring-platform/checker/internal/domain"

type SaveResultUseCase struct{}

func (s SaveResultUseCase) Save(result domain.CheckResult) bool {
	return result.Status == "down"
}