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
	v.SetDefault("color", "auto")
}

// bind command line flags to viper
func BindFlags(cmd *cobra.Command) {
	lookUp := cmd.Flags().Lookup
	v.BindPFlag("root_dir", lookUp("dir"))
	v.BindPFlag("color", lookUp("color"))
	v.BindPFlag("log_level", lookUp("log-level"))
}
