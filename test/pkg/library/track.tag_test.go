package library_test

import (
	"testing"

	"github.com/liampulles/banger/pkg/library"
	"github.com/stretchr/testify/assert"
)

func TestNewTagTrack_WhenFileDoesNotExist_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := "not.a.track"

	// Setup expectations
	expectedErr := "cannot create TagTrack - file error [not.a.track]: open not.a.track: no such file or directory"

	// Exercise SUT
	actual, err := library.NewTagTrack(fixture)

	// Verify results
	assert.Nil(t, actual)
	assert.EqualError(t, err, expectedErr)
}

func TestNewTagTrack_WhenFileIsNotASong_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := "testdata/not.a.song.txt"

	// Setup expectations
	expectedErr := "cannot create TagTrack - metadata error [testdata/not.a.song.txt]: seek testdata/not.a.song.txt: invalid argument"

	// Exercise SUT
	actual, err := library.NewTagTrack(fixture)

	// Verify results
	assert.Nil(t, actual)
	assert.EqualError(t, err, expectedErr)
}

func TestNewTagTrack_WhenFileIsASong_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := "testdata/actual.song.mp3"

	// Exercise SUT
	actual, err := library.NewTagTrack(fixture)

	// Verify results
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestTagTrack_Album_ShouldReturnAlbum(t *testing.T) {
	// Setup fixture
	path := "testdata/actual.song.mp3"
	fixture, err := library.NewTagTrack(path)
	assert.NoError(t, err)

	// Setup expectations
	expected := "WITCHY, BATTY, SPOOKY, HALLOWEEN IN SEPTEMBER !!"

	// Exercise SUT
	actual := fixture.Album()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestTagTrack_Artist_ShouldReturnArtist(t *testing.T) {
	// Setup fixture
	path := "testdata/actual.song.mp3"
	fixture, err := library.NewTagTrack(path)
	assert.NoError(t, err)

	// Setup expectations
	expected := "Loyalty Freak Music"

	// Exercise SUT
	actual := fixture.Artist()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestTagTrack_Genre_ShouldReturnGenre(t *testing.T) {
	// Setup fixture
	path := "testdata/actual.song.mp3"
	fixture, err := library.NewTagTrack(path)
	assert.NoError(t, err)

	// Setup expectations
	expected := "Ambient"

	// Exercise SUT
	actual := fixture.Genre()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestTagTrack_Location_ShouldReturnLocation(t *testing.T) {
	// Setup fixture
	path := "testdata/actual.song.mp3"
	fixture, err := library.NewTagTrack(path)
	assert.NoError(t, err)

	// Exercise SUT
	actual := fixture.Location()

	// Verify results
	assert.Equal(t, path, actual)
}

func TestTagTrack_Title_ShouldReturnTitle(t *testing.T) {
	// Setup fixture
	path := "testdata/actual.song.mp3"
	fixture, err := library.NewTagTrack(path)
	assert.NoError(t, err)

	// Setup expectations
	expected := "Ghost Surf Rock"

	// Exercise SUT
	actual := fixture.Title()

	// Verify results
	assert.Equal(t, expected, actual)
}
