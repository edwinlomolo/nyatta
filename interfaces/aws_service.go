package interfaces

import "mime/multipart"

type AwsServicesInterface interface {
	UploadFile(multipart.File, *multipart.FileHeader) (string, error)
}
