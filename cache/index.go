package cache

import (
	"os"

	"github.com/Yakiyo/tilde/utils"
	"github.com/Yakiyo/tilde/where"
	json "github.com/json-iterator/go"
)

// using string instead of `error` cz stupid ide screams when i use
// capitalized error strings. REEEEEEEEEEEEEEEEEEEEEE!

// read index.json from cache dir
func ReadIndex() (Index, string) {
	indexPath := where.Index()
	index := Index{}
	if !utils.FsExists(indexPath) {
		return index, "Cannot find index file in cache directory. Use the `--update/-u` flag to update cache"
	}
	f, err := os.ReadFile(indexPath)
	if err != nil {
		return index, err.Error()
	}
	err = json.Unmarshal(f, &index)
	return index, err.Error()
}

type Index struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	Name    string   `json:"name"`
	Targets []Target `json:"targets"`
}

type Target struct {
	Os       string `json:"os"`
	Language string `json:"language"`
}
