package middleware

import (
	"io"
	"log"
	"net/http"
	"time"
)

func MethodCheck(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, `{"error":"Method is not allowed"}`, http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Logging(internal_url string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		start := time.Now()
		next.ServeHTTP(w, req)
		body, _ := io.ReadAll(req.Body)
		if !(internal_url == req.RequestURI && http.MethodGet == req.Method && len(body) == 0) {
			log.Printf("%s %s %dus", req.Method, req.RequestURI, time.Since(start).Microseconds())
		}
	})
}
