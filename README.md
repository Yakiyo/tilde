# tilde
[![ci](https://github.com/Yakiyo/tilde/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/Yakiyo/tilde/actions/workflows/ci.yml) ![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Yakiyo/tilde) ![GitHub tag (with filter)](https://img.shields.io/github/v/tag/Yakiyo/tilde?label=version)


An easy to use and fast console client for [tldr](https://github.com/tldr-pages/tldr)

This was inspired by the amazing [tealdeer](https://github.com/dbrgn/tealdeer). My original goal was to see wether I could write something that has the same performance as tealdeer, but alas, I couldn't achive it. It just shows that tealdeer is that good! Nevertheless it was a fun project to do, and I managed to learn more Go.

## Installation

Download binaries from the [release](https://github.com/Yakiyo/tilde/releases/latest) section

## Usage
Tilde follows the [tldr client specifications](https://github.com/tldr-pages/tldr/blob/main/CLIENT-SPECIFICATION.md). The version of the spec followed can be viewed by running the `tldr --version` flag. 

Tilde stores a local cache of tldr pages in `~/.tilde/cache` directory. If it's missing, the app will warn you and you can add it with the `tldr --update` command

### Basics
```shell
 $ tldr git # search for git

 $ tldr git commit # search for git-log

 $ tldr --platform=windows cmd # search for `cmd` in windows

 $ tldr --language=bn zsh # search for `zsh` in Bangla language

 $ tldr -l cn -p osx tldr # combine `language` and `platform`

 $ tldr -f /path/to/some/file.md # render a custom markdown file

 $ tldr --raw gh # render without any formatting

```

You can customize the outputs with more flags such as `--color`, `--log-level` etc. 

The color flag takes one of `always` (always shows colors), `never` (disable colors) and `auto` (default). When set to auto, it tries to check if the output is a terminal and the env var `NO_COLOR` is not set.

The log-level flag can be used to customize log filters. Default is warn. You can set it to `info` to get some additional but not-necessary logs.

View the entire list by running `tldr --help`.

## Configuration
Tilde supports configuration via a [toml](https://toml.io) file. The default location of the file is `~/.tilde/tilde.toml`. This can specified using the command line `--config` flag. The config flag can be used to specify a custom config file to use. When the config flag is used, an error is thrown if the file does not exist in the mentioned path. If left unspecified, tilde tries to use the default file, if it is present, else uses defaults.

Supported config values
```toml
log_level = "warn" # set log level (one of info, warn, error, fatal)
color = "always" # set color (one of always, auto, never)
root_dir = "/path/to/tilde/root/dir" # tilde's root dir. Caches and other things are stored here. defaults to ~/.tilde
platform = "windows" # set default platform to use, default is user platform
language = "cn" # default language to use
custom_pages = "/path/to/dir" # set a dir containing custom pages. This are prioritized before the main cache files
```
The custom pages directory can be used to specify your own custom commands. This is given more priority. For example, if your custom directory contains a `zsh.md`, when using `tldr zsh`, the one in the custom directory is shown over the one in the tldr cache.

You can create a new config file by running `tldr --seed-config`. This creates a config file at `~/.tilde/tilde.toml`. **WARNING**: This will overwrite any existing configuration. `seed-config` uses the current config values. That is if used `tldr --log-level info --seed-config`, the newly created file will fill the `log_level` key with `info` instead of the default `warn`.
## Author

**tilde** © [Yakiyo](https://github.com/Yakiyo). Authored and maintained by Yakiyo.

Released under [MIT](https://opensource.org/licenses/MIT) License