package domain

import "time"

type MonitorCreated struct{ At time.Time }

func (e MonitorCreated) EventName() string     { return "monitors.created" }
func (e MonitorCreated) OccurredAt() time.Time { return e.At }

type MonitorPaused struct{ At time.Time }

func (e MonitorPaused) EventName() string     { return "monitors.paused" }
func (e MonitorPaused) OccurredAt() time.Time { return e.At }

type MonitorDeleted struct{ At time.Time }

func (e MonitorDeleted) EventName() string     { return "monitors.deleted" }
func (e MonitorDeleted) OccurredAt() time.Time { return e.At }
