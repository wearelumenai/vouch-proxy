package middleware

import (
	"net/http"

	"github.com/wearelumenai/clusauth/internal/conf"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// CORS log client requests
func CORS(conf conf.Conf) mux.MiddlewareFunc {
	var cors = cors.New(cors.Options{
		AllowedOrigins:   conf.Clusauth.Domains,
		Debug:            conf.Clusauth.Debug,
		AllowCredentials: true,
	})
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cors.HandlerFunc(w, r)
			if _, ok := w.Header()["Access-Control-Allow-Origin"]; ok {
				var origin = r.Header.Get("Origin")
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			next.ServeHTTP(w, r)
		})
	}
}
