package services

import (
	"context"

	cfg "github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/interfaces"
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

func NewAwsServices(cfg cfg.AwsConfig) *AwsServices {
	config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretAccessKey, "")),
	)
	if err != nil {
		log.Errorf("Unable to load aws config: %v", err)
	}

	// Create S3 client
	s3Client := manager.NewUploader(s3.NewFromConfig(config))

	return &AwsServices{S3: s3Client}
}

// UploadFile - upload file to s3
func (a *AwsServices) UploadFile() (string, error) { // TODO figure to input buffer file from js to here
	// Upload input params
	params := &s3.PutObjectInput{
		Bucket: &a.Config.S3.Buckets.Caretaker,
	}

	// Do upload
	res, err := a.S3.Upload(context.Background(), params)
	if err != nil {
		return "", err
	}
	return res.Location, nil
}
