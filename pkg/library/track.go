package library

// Track represents information about a track
type Track interface {
	Title() string
	Artist() string
	Album() string
	Genre() string
	Year() int
	Location() string
}
