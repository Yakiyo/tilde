package render

import (
	"strings"

	"github.com/Yakiyo/tilde/utils"
	"github.com/charmbracelet/log"
)

// base render function
func Render(file string, style string) {
	out, err := utils.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	style = strings.ToLower(style)
	switch style {
	case "raw":
		Raw(out)
	case "plain":
		Plain(out)
	default:
		if style != "fancy" {
			log.Warn("Invalid style value receive. Defaulting to `fancy`", "style", style)
		}
		Fancy(out)
	}
}
