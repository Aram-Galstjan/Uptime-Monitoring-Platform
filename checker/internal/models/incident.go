package models

import "time"

// Incident представляет инцидент с недоступностью сайта.
type Incident struct {
	SiteID    int
	StartedAt time.Time
	EndedAt   time.Time
	IsActive  bool
}
