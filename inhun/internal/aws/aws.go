package aws

import (
	"context"

	iconfig "github.com/inhun/GoropBox/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func LoadS3Client(cfg iconfig.AwsConfig) (*s3.Client, error) {
	awscfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")),
		config.WithRegion("ap-northeast-2"),
	)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(awscfg)

	return client, nil
}
