package application

import schedulerdomain "uptime-monitoring-platform/go-monolith/internal/scheduler/domain"

type Publisher interface {
	Publish(job schedulerdomain.CheckJob) error
}

func DispatchJob(publisher Publisher, job schedulerdomain.CheckJob) error {
	return publisher.Publish(job)
}
