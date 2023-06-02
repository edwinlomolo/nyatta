package config

type AwsConfig struct {
	AccessKey       string
	SecretAccessKey string
	S3              struct {
		Buckets struct {
			Caretaker string
		}
	}
}
