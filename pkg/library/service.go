package library

import (
	"fmt"

	"github.com/liampulles/banger/pkg/file"
)

type Service interface {
	PipeAllTracks() ([]Track, error)
}

type ServiceImpl struct {
	rootPath string
}

var _ Service = &ServiceImpl{}

func NewService(rootPath string) *ServiceImpl {
	return &ServiceImpl{
		rootPath: rootPath,
	}
}

func (s *ServiceImpl) PipeAllTracks() ([]Track, error) {
	paths, err := file.FindAllFilesRecursively(s.rootPath)
	if err != nil {
		return nil, fmt.Errorf("could not pipe tracks - find paths error: %w", err)
	}

	var tracks []Track
	for _, path := range paths {
		track, err := NewTagTrack(path)
		if err != nil {
			continue
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}
