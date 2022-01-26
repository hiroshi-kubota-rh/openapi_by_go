package test

import (
	"net/http"
)

// Handle health check request
func NewRouter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte("OK 2022/1/3"))
		if err != nil {
			panic(err)
		}
	}
}
