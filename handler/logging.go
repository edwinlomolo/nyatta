package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
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
		sqlStore := ctx.Value("sqlStore").(*store.Queries)

		var mpesaCallbackResponse *services.MpesaCallBackResponse

		if err := json.NewDecoder(r.Body).Decode(&mpesaCallbackResponse); err != nil {
			logger.Errorf("%s:%v", "MpesaChargeCallbackReadingBodyRequestError", err)
		}

		if mpesaCallbackResponse != nil && mpesaCallbackResponse.Body.StkCallBack.ResultCode == 0 {
			mpesaId := (mpesaCallbackResponse.Body.StkCallBack.CallBackMetadata.Item[1].Value).(string)
			amount := (mpesaCallbackResponse.Body.StkCallBack.CallBackMetadata.Item[0].Value).(float64)

			amountString := strconv.FormatFloat(amount, 'g', -1, 64)
			amountInt, err := strconv.Atoi(amountString)
			if err != nil {
				logger.Errorf("%s:%v", "MpesaChargeCallbackAmountParsing", err)
			}

			invoiceId := mpesaCallbackResponse.Body.StkCallBack.CheckoutRequestID

			updatedInvoice, err := sqlStore.UpdateInvoiceForMpesa(ctx, store.UpdateInvoiceForMpesaParams{
				MpesaID:       sql.NullString{String: mpesaId, Valid: true},
				WCoCheckoutID: sql.NullString{String: invoiceId, Valid: true},
				Status:        model.InvoiceStatusProcessed,
				Amount:        sql.NullInt32{Int32: int32(amountInt), Valid: true},
			})
			if err != nil {
				logger.Errorf("%s:%v", "MpesaChargeCallbackUpdateInvoiceError", err)
			}

			if _, err := sqlStore.UpdateLandlord(ctx, store.UpdateLandlordParams{IsLandlord: sql.NullBool{Bool: true, Valid: true}, Phone: updatedInvoice.Phone.String}); err != nil {
				logger.Errorf("%s:%v", "MpesaCharkCallbackUpdateLandlordStatusError", err)
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	})
}
