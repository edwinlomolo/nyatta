package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/3dw1nM0535/nyatta/services"
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

func MpesaChargeCallback() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := ctx.Value("log").(*logrus.Logger)

		var mpesaCallbackResponse *services.MpesaCallBackResponse

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Errorf("%s:%v", "MpesaCallbackError", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.Unmarshal(reqBody, &mpesaCallbackResponse)
		logger.Infoln(mpesaCallbackResponse)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	})
}
