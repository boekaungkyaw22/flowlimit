package flowlimit

import "time"

type Decision struct {
	Allowed bool

	RetryAfter time.Duration

	Remaining int

	ResetAt time.Time
}
