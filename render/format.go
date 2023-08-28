package render

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fatih/color"
	"github.com/samber/lo"
)

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
			line = line[1 : len(line)-1]
			line = highLightVariable(line)
			line = color.CyanString(line)
			res = append(res, "      "+line)
		}
	}
	res = append(res, "")
	return res
}

// FIXME: this needs work. Run `~/.tilde/cache/pages/common/git.md` to see the mistakes
// in the printing format

// parse an `example`, remove `{{`, `}}` and underline the variable
func highLightVariable(line string) string {
	lines := strings.Split(line, "}}")
	lines = lo.Map(lines, func(line string, _ int) string {
		l := strings.Split(line, "{{")
		if len(l) > 1 {
			l[1] = underline(l[1])
		}
		return strings.Join(l, "{{")
	})
	line = strings.Join(lines, "}}")
	line = strings.ReplaceAll(line, "{{", "")
	line = strings.ReplaceAll(line, "}}", "")
	return line
}
