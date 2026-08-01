package queue

import schedulerdomain "uptime-monitoring-platform/go-monolith/internal/scheduler/domain"

type RedisPublisher struct{}

func (p RedisPublisher) Publish(job schedulerdomain.CheckJob) error {
	_ = job
	return nil
}
