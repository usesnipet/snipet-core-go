package provider

import "io"

type Object struct {
	ObjectRefID string
	MimeType    string
	Size        int64

	Text       *string
	Binary     []byte
	Structured map[string]any
	Stream     io.ReadCloser
}
