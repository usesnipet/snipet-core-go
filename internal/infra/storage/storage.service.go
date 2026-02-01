package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/usesnipet/snipet-core-go/internal/config"
)

type Service struct {
	s3     *s3.Client
	bucket string
}

func NewStorageService(s3 *s3.Client) *Service {
	return &Service{
		s3:     s3,
		bucket: config.GetEnv().STORAGE_BUCKET,
	}
}

func (s *Service) Get(
	ctx context.Context,
	key string,
) (io.ReadCloser, error) {
	out, err := s.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	return out.Body, nil
}

func (s *Service) Delete(ctx context.Context, key string) error {
	_, err := s.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	return err
}
