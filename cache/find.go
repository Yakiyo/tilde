package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/tilde/utils"
	"github.com/Yakiyo/tilde/where"
	"github.com/charmbracelet/log"
	"github.com/mitchellh/go-homedir"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

func Find(command string) string {
	// check if there is any custom dir specified by user, and if it contains any file
	// with that command name
	if custom := viper.GetString("custom_dir"); custom != "" {
		custom, err := homedir.Expand(custom)
		if err != nil {
			log.Fatal("Unable to expand custom directory value", "value", custom, "err", err)
		}
		file := filepath.Join(custom, command+".md")
		if utils.FsExists(file) {
			return file
		}
	}
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
