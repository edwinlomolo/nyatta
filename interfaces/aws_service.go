package interfaces

import "github.com/99designs/gqlgen/graphql"

type AwsServicesInterface interface {
	UploadFile(graphql.Upload) (string, error)
}
