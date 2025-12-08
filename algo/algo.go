package algo

import "time"



type Result struct {
	NextState any
	Allowed  bool
	RetryAfter time.Duration
	Remaining int
	ResetAt time.Time
	ExpiresAt time.Time
}

type Algorithm interface {
	Evaluate(prevState any, now time.Time) Result
}