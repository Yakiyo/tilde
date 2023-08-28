package render

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fatih/color"
)

var variableRegex = regexp.MustCompile(`\{\{(.*)\}\}`)
var underline = color.New(color.Underline).Sprint

// format markdown
func format(md string) []string {
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
			res = append(res, color.HiGreenString(line[1:]))
		} else if strings.HasPrefix(line, "`") {
			// example
			line = highLightVariable(line[1 : len(line)-1])
			res = append(res, "      "+color.CyanString(line))
		}
	}
	res = append(res, "")
	return res
}

// parse an `example`, remove `{{`, `}}` and underline the variable
func highLightVariable(line string) string {
	return variableRegex.ReplaceAllString(line, underline("$1"))
}
