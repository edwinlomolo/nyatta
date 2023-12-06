package services

import (
	"bytes"
	"context"
	"mime/multipart"

	cfg "github.com/3dw1nM0535/nyatta/config"
	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
)

// AwsService client
type AwsService interface {
	UploadGqlFile(graphql.Upload) (string, error)
	UploadRestFile(multipart.File, *multipart.FileHeader) (string, error)
}

type awsClient struct {
	S3     *manager.Uploader
	Config cfg.AwsConfig
	log    *logrus.Logger
}

// NewAwsService - factory for aws service
func NewAwsService(cfg cfg.AwsConfig, logger *logrus.Logger) AwsService {
	config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretAccessKey, "")),
	)
	if err != nil {
		logger.Errorf("Unable to load aws config: %v", err)
	}

	// Create S3 client
	s3Client := manager.NewUploader(s3.NewFromConfig(config))

	return &awsClient{S3: s3Client, Config: cfg, log: logger}
}

// UploadGqlFile - graphql upload
func (a *awsClient) UploadGqlFile(file graphql.Upload) (string, error) {
	params := &s3.PutObjectInput{
		Bucket: aws.String(a.Config.S3.Buckets.Media),
		Key:    aws.String(file.Filename),
		Body:   file.File,
	}

	res, err := a.S3.Upload(context.Background(), params)
	if err != nil {
		a.log.Errorf("%s:%v", a.ServiceName(), err)
		return "", err
	}
	return res.Location, nil
}

// UploadRestFile - upload file to s3
func (a *awsClient) UploadRestFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	params := &s3.PutObjectInput{
		Bucket: aws.String(a.Config.S3.Buckets.Media),
		Key:    aws.String(fileHeader.Filename),
		Body:   bytes.NewReader(buffer),
	}

	// Do upload
	res, err := a.S3.Upload(context.Background(), params)
	if err != nil {
		a.log.Errorf("%s: %v", a.ServiceName(), err)
		return "", err
	}
	return res.Location, nil
}

// ServiceName - get service name
func (a *awsClient) ServiceName() string {
	return "awsClient"
}
