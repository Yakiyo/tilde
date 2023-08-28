package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/tilde/utils"
	"github.com/Yakiyo/tilde/where"
	"github.com/charmbracelet/log"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

// // find a command
// func Find(command string) Command {
// 	platform := viper.GetString("platform")
// 	language := viper.GetString("language")
// 	c, err := ReadIndex()
// 	if err != "" {
// 		log.Error(err)
// 		os.Exit(1)
// 	}
// 	// t := Target{
// 	// 	Os:       platform,
// 	// 	Language: language,
// 	// }
// 	f, b := lo.Find[Command](c.Commands, func(i Command) bool {
// 		return i.Name == command
// 	})
// 	if !b {
// 		fmt.Fprintf(
// 			os.Stderr,
// 			"Page `%v` not found in cache\nUpdate the cache with `tldr -u` or submit a pr via the following link:\n%v%v",
// 			command,
// 			`https://github.com/tldr-pages/tldr/issues/new?title=page%20request:%20`,
// 			command,
// 		)
// 	}
// 	if !lo.ContainsBy(f.Targets, func(t Target) bool { return t.Os == platform }) {
// 		log.Infof("Page %v not found for platform %v", command, platform)
// 	}
// 	// if platform was not the user platform, we try it first, else we skip too common
// 	if platform != utils.Platform() {}

// 	return f
// }

func Find(command string) string {
	language := viper.GetString("language")
	platform := utils.SafePlatform(viper.GetString("platform"))
	if !utils.IsValidPlatform(platform) {
		log.Error("Invalid value for platform. Must be one of "+strings.Join(utils.VALID_PLATFORMS, ", "), "platform", platform)
		os.Exit(1)
	}
	var lang_dir string = "pages"
	if language != "en" {
		lang_dir += "." + language
	}
	// TODO: check for custom pages dir here

	// create a sorted slice of platforms, based on priority
	// first comes user provided platform, then the default system platform
	// then common, and then the rest
	// we use a set, so that dupes get removed
	platforms := []string{platform}
	if platform != utils.Platform() {
		platforms = append(platforms, utils.Platform())
	}
	platforms = append(platforms, "common")
	for _, p := range utils.VALID_PLATFORMS {
		if !lo.Contains(platforms, p) {
			platforms = append(platforms, p)
		}
	}
	log.Debug("platform hierarchy determined", "slice", platforms)
	pages := filepath.Join(where.Cache(), lang_dir)
	log.Debug("Using pages dir", "dir", pages)

	cmd_file := fmt.Sprintf("%v.md", command)

	var path string
	for _, platform := range platforms {
		log.Info("Searching for page", "platform", platform)
		path = filepath.Join(pages, platform, cmd_file)
		// if file exists, use it
		if utils.FsExists(path) {
			return path
		}
	}
	return ""
}
