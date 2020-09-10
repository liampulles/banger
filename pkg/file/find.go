package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindAllFilesRecursively(rootPath string) ([]string, error) {
	if !DoesExist(rootPath) || !IsDir(rootPath) {
		return nil, fmt.Errorf("could not find all paths in tree - not a dir: %s", rootPath)
	}

	var paths []string
	var errors []error
	err := filepath.Walk(rootPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				errors = append(errors, err)
				return nil
			}
			if info.IsDir() {
				return nil
			}
			paths = append(paths, path)
			return nil
		})

	if err != nil {
		return nil, fmt.Errorf("could not find all paths in tree - walk error: %w", err)
	}

	if len(errors) > 0 {
		fmt.Fprintf(os.Stderr, "encountered %d errors walking tree - continuing\n", len(errors))
	}
	return paths, nil
}
