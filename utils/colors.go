package utils

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/mattn/go-isatty"
)

// wether stdout is terminal or not
func IsAtty() bool {
	fd := os.Stdout.Fd()
	return (isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd))
}

// Calculate wether colors should be enabled or not
func ColorOn(level string) bool {
	switch strings.ToLower(level) {
	case "never":
		return false
	case "always":
		return true
	case "auto":
		_, noColor := os.LookupEnv("NO_COLOR")
		return !noColor && IsAtty()
	}
	log.Warn("Invalid color level provided. Must be one of auto, always, never. Using 'always' by default")
	return true
}
