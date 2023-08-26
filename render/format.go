package render

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fatih/color"
)

// format markdown
func format(md string) []string {
	lines := strings.Split(md, "\n")
	res := []string{}
	res = append(res, "")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			log.Debug("Skipping command name")
			continue
		} else if strings.HasPrefix(line, ">") {
			// command description
			res[0] += color.BlueString(line) + "\n"
		} else if strings.HasPrefix(line, "-") {
			// example description
			res = append(res, line[1:])
		} else if strings.HasPrefix(line, "`") {
			// example
			line = line[1 : len(line)-1]
			res = append(res, "      "+line)
		}
	}
	return res
}
