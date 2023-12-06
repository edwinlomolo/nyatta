package handler

import (
	"encoding/json"
	"net/http"

	"github.com/3dw1nM0535/nyatta/services"
	"github.com/sirupsen/logrus"
)

func UploadHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := ctx.Value("log").(*logrus.Logger)
		maxSize := int64(6000000)

		err := r.ParseMultipartForm(maxSize)
		if err != nil {
			er := "FileTooLargeError"
			log.Errorf("%s:%v", err, err)
			http.Error(w, er, http.StatusBadRequest)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			er := "NoFileUploadedError"
			log.Errorf("%s:%v", er, err)
			http.Error(w, er, http.StatusBadRequest)
			return
		}
		defer file.Close()

		imageUri, err := ctx.Value("awsService").(services.AwsService).UploadRestFile(file, fileHeader)
		if err != nil {
			er := "AwsServicesInternalError"
			log.Errorf("%s:%v", er, err)
			http.Error(w, er, http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(struct {
			ImageUri string `json:"image_uri"`
		}{ImageUri: imageUri})
		if err != nil {
			er := "JsonMarshalError"
			log.Errorf("%s:%v", er, err)
			http.Error(w, er, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	})
}
