package utils

import (
	"runtime"
	"strings"

	"github.com/samber/lo"
)

// Get current platform name. Return common if unknown (shouldnt generally occur)
func Platform() string {
	switch runtime.GOOS {
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	case "osx", "macos", "darwin":
		return "osx"
	}
	return "common"
}

// Possible valid platforms
var VALID_PLATFORMS = []string{"windows", "linux", "osx", "android", "sunos"}

// check if a platform is valid or not
func IsValidPlatform(plat string) bool {
	return lo.Contains(VALID_PLATFORMS, plat)
}

// convert plat to app friendly type
func SafePlatform(plat string) string {
	plat = strings.ToLower(plat)
	switch plat {
	case "darwin", "macos":
		plat = "osx"
	}
	return plat
}
