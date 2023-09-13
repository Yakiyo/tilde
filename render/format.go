package render

import (
	"strings"

	"github.com/Yakiyo/tilde/utils"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

var reset = "\x1b[0m"
var green = "\x1b[32m"
var cyan = "\x1b[36m"
var cyanUnd = "\x1b[36;4m"

// format markdown
func format(md string) []string {
	setColor()
	lines := strings.Split(md, "\n")
	res := []string{}
	res = append(res, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			// command name
			log.Debug("Skipping command name")
			continue
		} else if strings.HasPrefix(line, ">") {
			// command description
			res[0] += line[1:] + "\n"
		} else if strings.HasPrefix(line, "-") {
			// example description
			res = append(res, desc(line[1:]))
		} else if strings.HasPrefix(line, "`") {
			// example
			line = line[1 : len(line)-1]
			line = example(line)
			res = append(res, "      "+line)
		}
	}
	res = append(res, "")
	return res
}

func setColor() {
	level := viper.GetString("color")
	if !utils.ColorOn(level) {
		log.Info("Switching off colors")
		reset = ""
		cyan = ""
		cyanUnd = ""
		green = ""
	}
}

func desc(s string) string { return green + s + reset }

func example(line string) string {
	line = cyan + line + reset
	line = strings.ReplaceAll(line, "{{", reset+cyanUnd)
	line = strings.ReplaceAll(line, "}}", reset+cyan)
	return line
}
