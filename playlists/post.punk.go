package playlists

import (
	"strings"

	"github.com/liampulles/banger/pkg/library"
)

func PostPunk(in []library.Track) []library.Track {
	collected := make([]library.Track, 0)
	for _, track := range in {
		genre := toRegular(track.Genre())
		if strings.Contains(genre, "POST") && strings.Contains(genre, "PUNK") {
			collected = append(collected, track)
		}
	}
	return collected
}
