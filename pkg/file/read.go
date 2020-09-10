package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadBytes(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file - open error: %w", err)
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read file - readAll error: %w", err)
	}

	return bytes, nil
}
