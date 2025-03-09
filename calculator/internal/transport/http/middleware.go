package router

import (
	"io"
	"log"
	"net/http"
)

func Logging(internal_url string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		body, _ := io.ReadAll(r.Body)
		if !(internal_url == r.RequestURI && http.MethodGet == r.Method && len(body) == 0) {
			log.Printf("%s %s", r.Method, r.RequestURI)
		}
	})
}
