package config

import (
	"github.com/Yakiyo/tilde/where"
	v "github.com/spf13/viper"
)

func init() {
	v.SetConfigType("toml")
	v.SetConfigName("tilde")

	// to be used when setting values from cli
	v.RegisterAlias("tilde_dir", "dir")

	v.SetDefault("auto_update", true)
	v.SetDefault("log_level", "warn")
	v.SetDefault("tilde_dir", where.Dir())
	// must be one of always, auto, never
	v.SetDefault("colors", "auto")
	// must be one of fancy (use glow), plain (selt formatter), none/raw (raw value)
	v.SetDefault("style", "fancy")
}
