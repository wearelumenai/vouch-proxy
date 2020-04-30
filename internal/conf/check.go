package conf

import (
	"fmt"
	"net/http"

	"errors"
)

// ErrNoPing raised if endpoint does not ping
var ErrNoPing = errors.New("no ping")

// Ping if input endpoint ping
func Ping(endpoint string, path string) (err error) {
	var url = fmt.Sprintf("%s/%s", endpoint, path)
	var resp *http.Response
	resp, err = http.Get(url)
	if err == nil {
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			err = ErrNoPing
		}
	}
	return
}
