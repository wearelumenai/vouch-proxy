package handler

import (
	"net/http"

	"github.com/wearelumenai/clusauth/internal/conf"
)

// Ping handler
func Ping(conf conf.Conf) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err = conf.Vouch.Ping()
		if err == nil {
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte("pong"))
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
