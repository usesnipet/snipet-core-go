package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	appConfig "github.com/usesnipet/snipet-core-go/internal/config"
)

func NewS3Client() (*s3.Client, error) {
	env := appConfig.GetEnv()

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(env.STORAGE_REGION),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				env.STORAGE_ACCESS_KEY,
				env.STORAGE_SECRET_KEY,
				"",
			),
		),
		config.WithBaseEndpoint(env.STORAGE_ENDPOINT),
	)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = env.STORAGE_PATH_STYLE
	}), nil
}
