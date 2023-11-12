package handler

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/3dw1nM0535/nyatta/services"
	"github.com/sirupsen/logrus"
)

func MpesaChargeCallback() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var paystackMpesaCallbackResponse *services.PaystackCallbackResponse

		ctx := r.Context()
		logger := ctx.Value("log").(*logrus.Logger)

		paystackSignature := r.Header.Get("x-paystack-signature")
		paystackSecretKey := os.Getenv("PAYSTACK_SECRET_KEY")

		hash := hmac.New(sha512.New, []byte(paystackSecretKey))

		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Errorf("%s:%v", "PaystackMpesaChargeCallbackReadingBodyRequestError", err)
		}

		hash.Write(body)
		expectedMac := hex.EncodeToString(hash.Sum(nil))
		if expectedMac != paystackSignature {
			logger.Errorf("%s:%v", "PaystackMpesaCallbackInvalidSignature", err)
			http.Error(w, "Invalid signature", http.StatusUnauthorized)
		}

		json.Unmarshal(body, &paystackMpesaCallbackResponse)

		if err := ctx.Value("paystackService").(*services.PaystackServices).ReconcilePaystackMpesaCallback(*paystackMpesaCallbackResponse); err != nil {
			logger.Errorf("%s:%v", "PaystackMpesaChargeCallbackReadingBodyRequestError", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	})
}
