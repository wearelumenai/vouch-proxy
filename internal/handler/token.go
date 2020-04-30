package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/pingcap/errors"
	"github.com/wearelumenai/clusauth/internal/conf"
)

// Claims structure
type Claims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

// ErrInvalid raised if token is invalid
var ErrInvalid = errors.New("Invalid token")

// ParseToken parse token
func ParseToken(secret string, signingString string) (user string, err error) {
	var claims Claims
	_, err = jwt.ParseWithClaims(signingString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err == nil {
		user = claims.User
	}
	return
}

// MakeToken function
func MakeToken(secret string, user string) (signingString string, err error) {
	var signingKey = []byte(secret)
	var claims = Claims{
		User: user,
	}
	var token = jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(signingKey)
}

// Token returns token
func Token(conf conf.Conf) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var statusCode, _, header, err = ProxyVouch(conf.Vouch, "/validate", r.Header)

		if err == nil {
			if statusCode == http.StatusOK {
				var token string
				var user = header.Get("X-Clusauth-User")
				token, err = MakeToken(conf.Clusauth.Secret, user)
				w.Write([]byte(token))
			} else {
				w.WriteHeader(statusCode)
			}
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
