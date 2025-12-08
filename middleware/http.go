// file: middleware/http.go
package middleware

import (
	"net/http"

	"github.com/boekaungkyaw22/flowlimit"
)

// HTTPMiddleware returns a standard net/http middleware that uses flowlimit.
func HTTPMiddleware(l *flowlimit.Limiter, keyFunc func(r *http.Request) string) func(next http.Handler) http.Handler {
	// keyFunc decides how to build the rate-limit key (IP, user, etc.)
	if keyFunc == nil {
		keyFunc = defaultKeyFunc
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := keyFunc(r)

			decision, err := l.Allow(r.Context(), key)
			if err != nil {
				// In a real lib, you might let users customize error handling.
				http.Error(w, "internal rate limit error", http.StatusInternalServerError)
				return
			}

			if !decision.Allowed {
				// Default rejection behavior. You will later allow customizing this.
				w.Header().Set("Retry-After", decision.RetryAfter.String())
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func defaultKeyFunc(r *http.Request) string {
	// Very naive default: limit by remote address.
	// In real setups you might want IP parsing, headers, auth info, etc.
	return r.RemoteAddr
}
