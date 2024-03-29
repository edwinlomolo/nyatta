package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type LoggingHandler struct{}

// Logging - feed custom logger onto any request handler through context
func (l *LoggingHandler) Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := ctx.Value("log").(*logrus.Logger)
		// Some info on what is happening with request(s)
		log.Infof("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto)
		// next
		h.ServeHTTP(w, r)
	})
}
