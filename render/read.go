package render

import (
	"log"
	"os"

	"github.com/Yakiyo/tilde/utils"
)

// read a file's content to string
func ReadFile(path string) string {
	if !utils.FsExists(path) {
		log.Fatal("No file exists at path", "path", path)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
} 