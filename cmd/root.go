// command line handling
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Yakiyo/tilde/cache"
	"github.com/Yakiyo/tilde/config"
	"github.com/Yakiyo/tilde/meta"
	"github.com/Yakiyo/tilde/render"
	"github.com/Yakiyo/tilde/utils"
	"github.com/Yakiyo/tilde/where"
	"github.com/charmbracelet/log"
	"github.com/fatih/color"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "tldr",
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
		where.SetRoot(viper.GetString("root_dir"))
		color.NoColor = !utils.ColorOn(viper.GetString("color"))
		return nil
	},
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Viper settings", "config", viper.AllSettings())

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
			return
		}

		if clear_cache := utils.Must(cmd.Flags().GetBool("clear-cache")); clear_cache {
			cache := where.Cache()
			if !utils.FsExists(cache) {
				fmt.Println("Local cache does not exist. Nothing to remove")
				return
			}
			err := os.RemoveAll(cache)
			if err != nil {
				log.Error("Failed to clear cache", "error", err)
			}
		}

		if list := utils.Must(cmd.Flags().GetBool("list")); list {
			cache.List()
		}

		if seed := utils.Must(cmd.Flags().GetBool("seed-config")); seed {
			dir := where.Dir()
			if !utils.FsExists(dir) {
				os.MkdirAll(dir, os.ModePerm)
			}
			if err := viper.SafeWriteConfig(); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Successfully seeded config at", where.Config())
		}

		raw := utils.Must(cmd.Flags().GetBool("raw"))
		if rnd := utils.Must(cmd.Flags().GetString("render")); rnd != "" {
			rnd = utils.Must(homedir.Expand(rnd))
			render.Render(rnd, raw)
			return
		}

		if len(args) < 1 {
			log.Error("No args provided. Must provided at least 1 argument")
			os.Exit(1)
		}

		c := strings.ToLower(strings.Join(args, "-"))
		f := cache.Find(c)
		if f == "" {
			fmt.Fprintf(
				os.Stderr,
				"%v Page `%v` not found in cache\nUpdate the cache with `tldr -u` or submit a pr via the following link:\n%v\n",
				color.RedString("ERROR:"),
				color.CyanString(c),
				color.HiCyanString(`https://github.com/tldr-pages/tldr/issues/new?title=page%20request:%20`+c),
			)
			os.Exit(1)
		}
		render.Render(f, raw)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
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
}
