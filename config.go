package flowlimit

import "time"

// Public Configuration for limiter algorithm
type AlgorithmType string

const (
	AlgorithmTokenBucket   AlgorithmType = "token_bucket"
	AlgorithmLeakyBucket   AlgorithmType = "leaky_bucket"
	AlgorithmSlidingWindow AlgorithmType = "sliding_window"
)

type Config struct {
	Limit int

	Window time.Duration

	Algorithm AlgorithmType

	Burst int // for Token Bucket (Optional, can be ignore for other algorithms)

	AlgorithmType AlgorithmType
}
