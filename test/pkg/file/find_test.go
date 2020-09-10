package file_test

import (
	"testing"

	"github.com/liampulles/banger/pkg/file"
	"github.com/stretchr/testify/assert"
)

func TestFindAllFilesRecursively_WhenFolderDoesNotExist_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := "not.a.folder"

	// Setup expectations
	expectedErr := "could not find all paths in tree - path error: stat not.a.folder: no such file or directory"

	// Exercise SUT
	actual, err := file.FindAllFilesRecursively(fixture)

	// Verify results
	assert.Nil(t, actual)
	assert.EqualError(t, err, expectedErr)
}

func TestFindAllFilesRecursively_WhenFolderDoesExist_ShouldReturnFiles(t *testing.T) {
	// Setup fixture
	fixture := "testdata"

	// Setup expectations
	expected := []string{
		"testdata/path1/subpath1/file.1.txt",
		"testdata/path1/subpath2/file.2.txt",
		"testdata/path1/subpath2/file.3.txt",
	}

	// Exercise SUT
	actual, err := file.FindAllFilesRecursively(fixture)

	// Verify results
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
