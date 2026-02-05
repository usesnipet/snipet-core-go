// Package storage provides an S3-backed object storage service for reading,
// writing, listing, and deleting objects in a configured bucket.
package storage

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/usesnipet/snipet-core-go/internal/config"
)

// ListEntry holds path, last modified time, and metadata for an object
// returned by List (S3 ListObjectsV2).
type ListEntry struct {
	Path         string     // Object key (path) in the bucket
	LastModified *time.Time // Last modification time
	Size         int64      // Size in bytes
	ETag         *string    // Entity tag for the object
	StorageClass string     // Storage class (e.g. STANDARD)
	OwnerID      *string    // Owner account ID
	OwnerDisplay *string    // Owner display name
}

type FileOutput struct {
	Metadata    map[string]string // The file metadata
	ContentType *string           // File content type (image/png, video/mp4, etc)
	Size        *int64            // The file size um bytes
}

// Service is an S3-backed storage service bound to a single bucket.
type Service struct {
	s3     *s3.Client
	bucket string
}

// NewStorageService builds a Service using the given S3 client and the
// bucket name from config (STORAGE_BUCKET).
func NewStorageService(s3 *s3.Client) *Service {
	return &Service{
		s3:     s3,
		bucket: config.GetEnv().STORAGE_BUCKET,
	}
}

// Get returns a reader for the object at the given key. The caller must close
// the returned io.ReadCloser when done.
func (s *Service) Get(
	ctx context.Context,
	key string,
) (io.ReadCloser, *FileOutput, error) {
	out, err := s.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})

	if err != nil {
		return nil, nil, err
	}
	output := FileOutput{
		Metadata:    out.Metadata,
		ContentType: out.ContentType,
		Size:        out.ContentLength,
	}
	return out.Body, &output, nil
}

// Delete removes the object at the given key from the bucket.
func (s *Service) Delete(ctx context.Context, key string) error {
	_, err := s.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	return err
}

// List returns all objects in the bucket. prefix may be nil to list the
// entire bucket, or a pointer to a prefix string to filter by key prefix.
// Results include path, last modified time, and metadata (size, ETag,
// storage class, owner). Pagination is handled automatically.
func (s *Service) List(ctx context.Context, prefix *string) ([]ListEntry, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: &s.bucket,
	}
	if prefix != nil && *prefix != "" {
		input.Prefix = prefix
	}

	paginator := s3.NewListObjectsV2Paginator(s.s3, input)
	var entries []ListEntry

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, obj := range page.Contents {
			entries = append(entries, listEntryFromS3(obj))
		}
	}

	return entries, nil
}

func listEntryFromS3(obj types.Object) ListEntry {
	e := ListEntry{
		Path:         ptrToString(obj.Key),
		LastModified: obj.LastModified,
		Size:         ptrToInt64(obj.Size),
		ETag:         obj.ETag,
		StorageClass: string(obj.StorageClass),
	}
	if obj.Owner != nil {
		e.OwnerID = obj.Owner.ID
		e.OwnerDisplay = obj.Owner.DisplayName
	}
	return e
}

func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ptrToInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}
