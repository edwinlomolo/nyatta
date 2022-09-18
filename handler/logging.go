package handler

import (
	"net/http"

	"go.uber.org/zap"
)

type LoggingHandler struct{}

func (l *LoggingHandler) Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := ctx.Value("log").(*zap.SugaredLogger)
		// Some info on what is happening with request(s)
		log.Infof("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto)
		// TODO: log request body contents?
		h.ServeHTTP(w, r)
	})
}
