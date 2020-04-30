package handler

import (
	"fmt"
	"net/http"

	"github.com/wearelumenai/clusauth/internal/conf"
	"github.com/wearelumenai/clusauth/internal/handler/middleware"

	"github.com/gorilla/mux"
)

// Handler return rounting handler
func Handler(conf conf.Conf) http.Handler {
	var router = mux.NewRouter()

	var vouch = Vouch(conf)

	for _, route := range []string{"login", "logout", "auth", "validate"} {
		var path = fmt.Sprintf("/%s", route)
		router.Schemes("https", "http").Methods(http.MethodGet).Path(path).Handler(vouch)
	}

	router.Schemes("http", "https").Methods(http.MethodGet).Path("/ping").HandlerFunc(Ping(conf))
	router.Schemes("http", "https").Methods(http.MethodGet).Path("/token").HandlerFunc(Token(conf))

	// apply middleware
	router.Use(
		middleware.Recover,
		middleware.Logger,
		middleware.CORS(conf),
	)

	return router
}
