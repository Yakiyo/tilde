package config

import (
	"github.com/Yakiyo/tilde/utils"
	"github.com/Yakiyo/tilde/where"
	"github.com/spf13/cobra"
	v "github.com/spf13/viper"
)

func init() {
	v.SetConfigType("toml")
	v.SetConfigName("tilde")
	v.AddConfigPath(where.Dir())

	// v.SetDefault("auto_update", true)
	v.SetDefault("log_level", "warn")
	v.SetDefault("root_dir", where.Dir())
	v.SetDefault("platform", utils.Platform())
	v.SetDefault("language", "en")
	v.SetDefault("custom_pages", "")
	// must be one of always, auto, never
	v.SetDefault("color", "auto")
}

// bind command line flags to viper
func BindFlags(cmd *cobra.Command) {
	lookUp := cmd.Flags().Lookup
	v.BindPFlag("root_dir", lookUp("dir"))
	v.BindPFlag("color", lookUp("color"))
	v.BindPFlag("log_level", lookUp("log-level"))
	v.BindPFlag("language", lookUp("language"))
	v.BindPFlag("platform", lookUp("platform"))
}
