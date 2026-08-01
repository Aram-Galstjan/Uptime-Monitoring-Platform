package domain

import "time"

type UserCreated struct{ At time.Time }

func (e UserCreated) EventName() string     { return "iam.user_created" }
func (e UserCreated) OccurredAt() time.Time { return e.At }

type OrgCreated struct{ At time.Time }

func (e OrgCreated) EventName() string     { return "iam.org_created" }
func (e OrgCreated) OccurredAt() time.Time { return e.At }
