package storage

import (
	"bytes"
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type PutOptions struct {
	ctx         context.Context
	key         string
	contentType string
	data        []byte

	// extra
	temp bool
	ttl  *time.Duration
}
type WithPutOption func(*PutOptions)

func defaultPutOptions(key string, data []byte) *PutOptions {
	return &PutOptions{
		ctx:         context.Background(),
		key:         key,
		data:        data,
		contentType: "application/octet-stream",
	}
}
func WithContext(ctx context.Context) WithPutOption {
	return func(o *PutOptions) {
		o.ctx = ctx
	}
}
func WithContentType(ct string) WithPutOption {
	return func(o *PutOptions) {
		o.contentType = ct
	}
}
func WithTemp() WithPutOption {
	return func(o *PutOptions) {
		o.temp = true
	}
}
func WithTTL(d time.Duration) WithPutOption {
	return func(o *PutOptions) {
		o.ttl = &d
	}
}
func (s *Service) Put(
	key string,
	data []byte,
	opts ...WithPutOption,
) error {
	o := defaultPutOptions(key, data)

	for _, opt := range opts {
		opt(o)
	}

	_, err := s.s3.PutObject(o.ctx, &s3.PutObjectInput{
		Bucket:      &s.bucket,
		Key:         &key,
		Body:        bytes.NewReader(data),
		ContentType: &o.contentType,
	})
	return err
}
