package utils

import (
	"runtime"
	"strings"
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
func ValidPlatform(plat string) bool {
	plat = strings.ToLower(plat)
	switch plat {
	case "darwin", "macos":
		plat = "osx"
	}
	for _, x := range VALID_PLATFORMS {
		if plat == x {
			return true
		}
	}
	return false
}
