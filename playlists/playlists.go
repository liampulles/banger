package playlists

import (
	"regexp"
	"strings"

	"github.com/liampulles/banger/pkg/library"
)

var Registered map[string]PlaylistCollector = map[string]PlaylistCollector{
	"post-punk":            PostPunk,
	"dummy":                Dummy,
	"nineties-alternative": NinetiesAlternative,
}

var alphanumeric, _ = regexp.Compile("[^a-zA-Z0-9]+")

type PlaylistCollector func([]library.Track) []library.Track

func toRegular(in string) string {
	regular := alphanumeric.ReplaceAllString(in, "")
	return strings.ToUpper(regular)
}
