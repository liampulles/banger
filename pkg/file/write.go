package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func WriteBytes(location string, bytes []byte) error {
	if err := mkPathsFor(location); err != nil {
		return err
	}
	if err := ioutil.WriteFile(location, bytes, 0666); err != nil {
		return fmt.Errorf("could not write to file - writeFile error: %w", err)
	}
	return nil
}

func mkPathsFor(fileLocation string) error {
	dirLocation, _ := path.Split(fileLocation)
	if err := os.MkdirAll(dirLocation, 0777); err != nil {
		return fmt.Errorf("could not make paths for file - mkdirAll error: %w", err)
	}
	return nil
}
