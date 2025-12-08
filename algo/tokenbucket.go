package algo

import (
	"time"
)

type TokenBucket struct {
	limit  int
	burst  int
	window time.Duration
}

func NewTokenBucket(limit int, window time.Duration, burst int) *TokenBucket {
	// If burst is not specified (0), default to limit or some reasonable value.
	// For now, let's assume if 0 it might just match limit or be handled by logic.
	// But simply setting it here.
	if burst <= 0 {
		burst = limit
	}
	return &TokenBucket{
		limit:  limit,
		window: window,
		burst:  burst,
	}
}

func (tb *TokenBucket) Evaluate(prevState any, now time.Time) Result {
	// Implementation placeholder - just returning a valid Result to satisfy interface
	// Real logic would go here
	return Result{
		Allowed:   true,
		Remaining: tb.limit,
		ResetAt:   now.Add(tb.window),
	}
}
