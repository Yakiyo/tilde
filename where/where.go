package where

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	homedir "github.com/mitchellh/go-homedir"
)

var root string

func init() {
	// set root to default location
	root = filepath.Join(home(), ".tilde")
}

// find home dir
func home() string {
	path, err := homedir.Dir()
	if err != nil {
		log.Fatal("Unable to locate user home dir, consider manually setting the value of $HOME/$USERPROFILE env var")
	}
	return path
}

// find root dir for use in tilde. This is `~/.tilde`
// maybe allow user to set it via command line flags or something in the future
func Dir() string {
	return root
}

// the file where the downloaded archive file will be kept before unpacking and finally being deleted afterwards
func Zip() string {
	return filepath.Join(os.TempDir(), "tldr.archive.zip")
}

// the dir where pages will be actually kept
func Cache() string {
	return filepath.Join(Dir(), "cache")
}

// path to index.json file located inside cache
func Index() string {
	return filepath.Join(Cache(), "index.json")
}

// set root to a path
func SetRoot(path string) {
	root = path
}
