package file

import (
	"path/filepath"
	"time"

	"github.com/usesnipet/snipet-core-go/internal/infra/provider"
)

func mapObject(
	providerName string,
	sourceID string,
	key string,
	size int64,
	lastModified time.Time,
) provider.ObjectRef {

	return provider.ObjectRef{
		ID:             sourceID + ":" + key,
		ObjectType:     detectType(key),
		SourceProvider: providerName,
		SourceID:       sourceID,
		Path:           key,
		Metadata: map[string]any{
			"size": size,
		},
		CreatedAt: lastModified,
	}
}

func detectType(key string) provider.ObjectType {
	ext := filepath.Ext(key)

	switch ext {
	case ".mp3", ".wav":
		return provider.ObjectAudio
	case ".png", ".jpg", ".jpeg":
		return provider.ObjectImage
	case ".mp4":
		return provider.ObjectVideo
	case ".json", ".csv":
		return provider.ObjectStructured
	default:
		return provider.ObjectFile
	}
}
