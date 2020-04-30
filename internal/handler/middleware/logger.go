package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger log client requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var start = time.Now()
		next.ServeHTTP(w, r)
		log.Printf("<< %s - %s %s %v (%v)", r.RemoteAddr, r.Method, r.URL, start, time.Since(start))
	})
}
