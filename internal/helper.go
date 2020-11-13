package internal

import "net/http"

func requestDriver(r *http.Request) (int64, error) {
	// token := r.URL.Query()["token"]
	// todo request driver out of token
	return int64(1), nil
}
