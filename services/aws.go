package services

import (
	"context"

	cfg "github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	log "github.com/sirupsen/logrus"
)

var _ interfaces.AwsServicesInterface = &AwsServices{}

type AwsServices struct {
	S3     *manager.Uploader
	Config cfg.AwsConfig
}

func NewAwsService(cfg cfg.AwsConfig) *AwsServices {
	config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretAccessKey, "")),
	)
	if err != nil {
		log.Errorf("Unable to load aws config: %v", err)
	}

	// Create S3 client
	s3Client := manager.NewUploader(s3.NewFromConfig(config))

	return &AwsServices{S3: s3Client, Config: cfg}
}

// UploadFile - upload file to s3
func (a *AwsServices) UploadFile(file graphql.Upload) (string, error) {
	// Upload input params
	params := &s3.PutObjectInput{
		Bucket: aws.String(a.Config.S3.Buckets.Caretaker),
		Key:    aws.String(file.Filename),
		Body:   file.File,
	}

	// Do upload
	res, err := a.S3.Upload(context.Background(), params)
	if err != nil {
		return "", err
	}
	return res.Location, nil
}
