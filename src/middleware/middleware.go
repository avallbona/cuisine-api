package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the request details along with the datetime
		fmt.Printf(
			"[%s] %s %s %s %s\n",
			time.Now().Format("2006-01-02 15:04:05 UTC"),
			time.Since(start),
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)
	})
}
