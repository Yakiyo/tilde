package config

import (
	"github.com/Yakiyo/tilde/where"
	"github.com/spf13/cobra"
	v "github.com/spf13/viper"
)

func init() {
	v.SetConfigType("toml")
	v.SetConfigName("tilde")
	v.AddConfigPath(where.Dir())

	v.SetDefault("auto_update", true)
	v.SetDefault("log_level", "warn")
	v.SetDefault("root_dir", where.Dir())
	// must be one of always, auto, never
	v.SetDefault("colors", "auto")
	// must be one of fancy (use glow), plain (selt formatter), none/raw (raw value)
	v.SetDefault("style", "fancy")
}

// bind command line flags to viper
func BindFlags(cmd *cobra.Command) {
	lookUp := cmd.Flags().Lookup
	v.BindPFlag("colors", lookUp("style"))
	v.BindPFlag("root_dir", lookUp("dir"))
}
