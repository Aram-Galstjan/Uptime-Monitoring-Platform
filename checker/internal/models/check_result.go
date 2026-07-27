package models

import "time"

// CheckResult представляет результат проверки сайта.
type CheckResult struct {
	SiteID       int
	Status       string
	ResponseTime int
	Error        string
	CheckedAt    time.Time
}
