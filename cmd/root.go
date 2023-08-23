// command line handling
package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/tilde/cache"
	"github.com/Yakiyo/tilde/config"
	"github.com/Yakiyo/tilde/meta"
	"github.com/Yakiyo/tilde/utils"
	"github.com/Yakiyo/tilde/where"
	"github.com/charmbracelet/log"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "tilde",
	Example: "tldr git log",
	Short:   "Fast tldr console client",
	Long: `tilde is a fast and frictionless console client for tldr.
	
View community driven and simplified man pages in your terminal`,
	Version: meta.Version,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// bind flags
		config.BindFlags(cmd)
		// set config file
		config.SetFile(utils.Must(cmd.Flags().GetString("config")))
		// read em
		config.Read()

		log.SetLevel(log.ParseLevel(viper.GetString("log_level")))
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug(viper.AllSettings())

		if len(os.Args[1:]) < 1 {
			cmd.Help()
			os.Exit(1)
		}
		if update := utils.Must(cmd.Flags().GetBool("update")); update {
			err := cache.Download()
			if err != nil {
				log.Fatal("Error downloading cache", "error", err)
			}
			fmt.Println("Successfully downloaded local cache")
		}

		if seed := utils.Must(cmd.Flags().GetBool("seed-config")); seed {
			if err := viper.WriteConfig(); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Successfully seeded config at", where.Config())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Bold,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDataType:   cc.Italic + cc.HiBlue,
		NoExtraNewlines: true,
	})

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(func() string {
		return `{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}` +
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
	f.StringP("style", "s", "", "Set output style [fancy, plain, raw]")
	f.StringP("language", "L", "", "Override language")
	f.StringP("platform", "p", utils.Platform(), "Override operating system")
}
