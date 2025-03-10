package router

import (
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		if !(strings.Contains(r.RequestURI, "internal")) {
			log.Infof("%s %s %dms", r.Method, r.RequestURI, time.Since(start).Milliseconds())
		}
	})
}
