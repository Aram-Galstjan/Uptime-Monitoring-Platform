package domain

import "time"

type CheckResult struct {
	Status  string
	Code    int
	Latency time.Duration
	Error   string
}