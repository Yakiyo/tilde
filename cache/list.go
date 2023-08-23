package cache

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	// "io/fs"
	// "path/filepath"

	"github.com/Yakiyo/tilde/utils"
	"github.com/Yakiyo/tilde/where"
	"github.com/charmbracelet/log"
)

func List() {
	if !utils.FsExists(where.Cache()) {
		log.Fatal("Page cache not found. Please download the cache by using the `--update/-u` flag\n\n" +
			"You can optionally enable auto cache downloading by adding `auto_update = true` in your config file\n" +
			"Path to config file: " + where.Config())
	}
	cache := filepath.Join(where.Cache(), "pages")
	m := map[string]bool{}
	dirs := []string{
		filepath.Join(cache, utils.Platform()),
		filepath.Join(cache, "common"),
	}
	fmt.Println(dirs)
	for _, dir := range dirs {
		filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				m[strings.TrimSuffix(d.Name(), ".md")] = true
			}
			return nil
		})
	}
	for k := range m {
		fmt.Println(k)
	}
}
