package playlists

import (
	"strings"

	"github.com/liampulles/banger/pkg/library"
)

func NinetiesAlternative(in []library.Track) []library.Track {
	collected := make([]library.Track, 0)
	for _, track := range in {
		year := track.Year()
		genre := toRegular(track.Genre())
		validGenre := strings.Contains(genre, "ALTERNATIVE") || strings.Contains(genre, "ELECTRONIC")
		if year >= 1990 && year < 2000 && validGenre {
			collected = append(collected, track)
		}
	}
	return collected
}
