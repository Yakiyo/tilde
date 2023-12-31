package cmd

import (
	"fmt"

	"github.com/Yakiyo/tilde/meta"
)

func init() {
	rootCmd.SetVersionTemplate(func() string {
		return `tilde {{printf "version %s" .Version}}` +
			fmt.Sprintf("\ntldr spec version %v\n", meta.TldrSpec)
	}())

	f := rootCmd.Flags()
	// config flags
	f.String("log-level", "", "Set log level [debug, info, warn, error]")
	f.StringP("dir", "d", "", "Set root directory to use for tilde")
	f.StringP("config", "c", "", "Set path to config file")
	f.String("color", "", "Enable or disable color output")

	// command-ish flags
	f.BoolP("list", "l", false, "List all commands in cache")
	f.BoolP("update", "u", false, "Update local cache")
	f.Bool("clear-cache", false, "Clear local cache")
	f.Bool("seed-config", false, "Creates the default configuration file at the default location")
	f.StringP("render", "f", "", "Render a local file")

	// command related flags
	f.BoolP("raw", "r", false, "Print raw markdown or not")
	f.StringP("language", "L", "", "Override language (Defaults to \"en\")")
	f.StringP("platform", "p", "", "Override operating system (Defaults to user platform)")
}
