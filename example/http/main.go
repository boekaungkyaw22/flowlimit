package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/boekaungkyaw22/flowlimit"
	"github.com/boekaungkyaw22/flowlimit/middleware"
)

func main() {
	// 1. Create a limiter with config
	limiter, err := flowlimit.NewLimiter(flowlimit.Config{
		Limit:     5,
		Window:    10 * time.Second,
		Burst:     5,
		Algorithm: flowlimit.AlgorithmTokenBucket,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 2. Create middleware, with key extractor
	rateLimitMiddleware := middleware.HTTPMiddleware(
		limiter,
		func(r *http.Request) string {
			// Simple IP-based strategy
			return "ip:" + r.RemoteAddr
		},
	)

	// 3. Example handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Request allowed! Time: %s\n", time.Now().Format(time.RFC3339))
	})

	// 4. Attach middleware to handler
	http.Handle("/", rateLimitMiddleware(handler))

	// 5. Start server
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
