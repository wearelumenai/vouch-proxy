package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/wearelumenai/clusauth/internal/conf"
)

// ProxyVouch generate a vouch request
func ProxyVouch(conf conf.Vouch, path string, iheader http.Header) (statusCode int, body []byte, header http.Header, err error) {
	var url = fmt.Sprintf("%s/%s", conf.Endpoint, path)
	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err == nil {
		for key, values := range iheader {
			req.Header[key] = values
		}
		req.Header.Add("X-Forwaded-Host", iheader.Get("Host"))
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err == nil {
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			header = resp.Header
			statusCode = resp.StatusCode
		}
	}

	return
}

// Vouch handler
func Vouch(conf conf.Conf) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/validate" {
			var authorization = r.Header.Get("Authorization")
			if authorization != "" {
				var authElts = strings.Split(authorization, " ")
				if len(authElts) == 2 && authElts[0] == "Bearer" {
					var signingString = authElts[1]
					var user, err = ParseToken(conf.Clusauth.Secret, signingString)
					if err == nil {
						w.Header().Set("X-Clusauth-User", user)
						return
					}
				}
			}
		}

		var statusCode, body, header, err = ProxyVouch(conf.Vouch, r.URL.Path, r.Header)
		if err == nil {
			for key, value := range header {
				header[key] = value
			}
			w.WriteHeader(statusCode)
			_, err = w.Write(body)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
