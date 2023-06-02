package interfaces

type AwsServicesInterface interface {
	UploadFile() (string, error)
}
