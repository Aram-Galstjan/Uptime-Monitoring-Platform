package domain

import "time"

type IncidentOpened struct{ At time.Time }

func (e IncidentOpened) EventName() string     { return "incidents.opened" }
func (e IncidentOpened) OccurredAt() time.Time { return e.At }

type IncidentResolved struct{ At time.Time }

func (e IncidentResolved) EventName() string     { return "incidents.resolved" }
func (e IncidentResolved) OccurredAt() time.Time { return e.At }

type IncidentEscalated struct{ At time.Time }

func (e IncidentEscalated) EventName() string     { return "incidents.escalated" }
func (e IncidentEscalated) OccurredAt() time.Time { return e.At }
