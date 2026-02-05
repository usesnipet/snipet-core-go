package provider

import "time"

type ObjectType string

const (
	ObjectText       ObjectType = "text"
	ObjectFile       ObjectType = "file"
	ObjectAudio      ObjectType = "audio"
	ObjectImage      ObjectType = "image"
	ObjectVideo      ObjectType = "video"
	ObjectStructured ObjectType = "structured"
	ObjectEvent      ObjectType = "event"
)

type ObjectRef struct {
	ID string

	ObjectType     ObjectType
	SourceProvider string // the name of the provider that created the object
	SourceID       string // the id of the source that created the object
	Path           string // the path of the object in the source

	Metadata  map[string]any
	CreatedAt time.Time
	UpdatedAt *time.Time
}
