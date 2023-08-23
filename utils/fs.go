package utils

import (
	"fmt"
	"os"
)

// exists returns whether the given file or directory exists
func FsExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// read a file's content to string
func ReadFile(path string) (string, error) {
	if !FsExists(path) {
		return "", fmt.Errorf("no file exists at path %v", path)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
