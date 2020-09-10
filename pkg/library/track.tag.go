package library

import (
	"fmt"
	"os"

	"github.com/dhowden/tag"
)

type TagTrack struct {
	path     string
	metadata tag.Metadata
}

var _ Track = &TagTrack{}

func NewTagTrack(path string) (*TagTrack, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot create TagTrack - file error [%s]: %w", path, err)
	}
	metadata, err := tag.ReadFrom(f)
	if err != nil {
		return nil, fmt.Errorf("cannot create TagTrack - metadata error [%s]: %w", path, err)
	}
	return &TagTrack{
		path:     path,
		metadata: metadata,
	}, nil
}

func (t *TagTrack) Album() string {
	return t.metadata.Album()
}

func (t *TagTrack) Artist() string {
	return t.metadata.Artist()
}

func (t *TagTrack) Genre() string {
	return t.metadata.Genre()
}

func (t *TagTrack) Location() string {
	return t.path
}

func (t *TagTrack) Title() string {
	return t.metadata.Title()
}

func (t *TagTrack) Year() int {
	return t.metadata.Year()
}
