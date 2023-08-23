package config

import (
	"github.com/Yakiyo/tilde/where"
	v "github.com/spf13/viper"
)

func init() {
	v.SetConfigType("toml")
	v.SetConfigName("tilde.toml")

	// to be used when setting values from cli
	v.RegisterAlias("tilde_dir", "dir")

	v.SetDefault("auto_update", true)
	v.SetDefault("log_level", "warn")
	v.SetDefault("tilde_dir", where.Dir())
	v.SetDefault("colors", "auto")
}
