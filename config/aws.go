package config

type AwsConfig struct {
	AccessKey       string `json:"accessKey"`
	SecretAccessKey string `json:"secretAccessKey"`
	S3              struct {
		Buckets struct {
			Media string `json:"media"`
		}
	}
}
