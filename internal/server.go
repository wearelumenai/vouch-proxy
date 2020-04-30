package internal

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/wearelumenai/clusauth/internal/conf"
	"github.com/wearelumenai/clusauth/internal/handler"

	"golang.org/x/crypto/acme/autocert"
)

// Serve deploy server
func Serve(conf conf.Conf) (err error) {
	var url *url.URL
	url, err = url.Parse(conf.Clusauth.Endpoint)

	var handler = handler.Handler(conf)

	var server = &http.Server{
		Addr:         fmt.Sprintf(":%v", url.Port()),
		ReadTimeout:  time.Duration(conf.Clusauth.Timeout) * time.Second,
		WriteTimeout: time.Duration(conf.Clusauth.Timeout) * time.Second,
		IdleTimeout:  time.Duration(conf.Clusauth.Timeout) * time.Second,
		Handler:      handler,
	}

	log.Printf("Listening on port %v", url.Port())

	if url.Scheme == "https" {
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("www.checknu.de"),
			Cache:      autocert.DirCache("/home/letsencrypt/"),
		}
		server.TLSConfig = &tls.Config{
			GetCertificate: m.GetCertificate,
		}
		err = server.ListenAndServeTLS(conf.Clusauth.Certfile, conf.Clusauth.Keyfile)
	} else {
		err = server.ListenAndServe()
	}

	return
}
