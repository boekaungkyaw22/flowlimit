package flowlimit

import (
	"context"
	"errors"

	"github.com/boekaungkyaw22/flowlimit/algo"
	"github.com/boekaungkyaw22/flowlimit/internal/timeutil"
	"github.com/boekaungkyaw22/flowlimit/storage"
)

type Limiter struct {
	cfg    Config
	engine algo.Algorithm
	store  storage.Store
	clock  timeutil.Clock
}

type Option func(*Limiter)

func WithClock(clock timeutil.Clock) Option {
	return func(l *Limiter) {
		l.clock = clock
	}
}

func WithStore(store storage.Store) Option {
	return func(l *Limiter) {
		l.store = store
	}
}

func NewLimiter(cfg Config, opts ...Option) (*Limiter, error) {
	algImpl, err := createAlgorithm(cfg)
	if err != nil {
		return nil, err
	}

	l := &Limiter{
		cfg:    cfg,
		engine: algImpl,
		store:  nil,
		clock:  &timeutil.SystemClock{},
	}

	for _, opt := range opts {
		opt(l)
	}

	return l, nil
}

func createAlgorithm(cfg Config) (algo.Algorithm, error) {
	switch cfg.Algorithm {
	case AlgorithmTokenBucket:
		// Fix: Pass individual parameters to avoid circular dependency
		return algo.NewTokenBucket(cfg.Limit, cfg.Window, cfg.Burst), nil
	case AlgorithmLeakyBucket:
		// return algo.NewLeakyBucket(cfg), nil
		return nil, errors.New("leaky bucket not implemented")
	case AlgorithmSlidingWindow:
		// return algo.NewSlidingWindow(cfg), nil
		return nil, errors.New("sliding window not implemented")
	default:
		return nil, ErrUnknownAlgorithm
	}
}

var ErrUnknownAlgorithm = errors.New("unknown algorithm")

func (l *Limiter) Allow(ctx context.Context, key string) (Decision, error) {
	now := l.clock.Now()
	entry, ok := l.store.Get(key)

	var prevState any
	if ok {
		prevState = entry.State
	}

	result := l.engine.Evaluate(prevState, now)

	newEntry := storage.Entry{
		State:     result.NextState,
		ExpiresAt: result.ExpiresAt,
	}

	l.store.Set(key, newEntry)

	return Decision{
		Allowed:    result.Allowed,
		RetryAfter: result.RetryAfter,
		Remaining:  result.Remaining,
		ResetAt:    result.ResetAt,
	}, nil
}
