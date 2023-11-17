package interfaces

import (
	"mime/multipart"

	"github.com/99designs/gqlgen/graphql"
)

type AwsServicesInterface interface {
	UploadGqlFile(graphql.Upload) (string, error)
	UploadRestFile(multipart.File, *multipart.FileHeader) (string, error)
}
