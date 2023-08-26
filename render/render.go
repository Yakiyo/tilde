package render

import (
	"fmt"
	"strings"

	"github.com/Yakiyo/tilde/utils"
	"github.com/charmbracelet/log"
)

// base render function
func Render(file string, raw bool) {
	out, err := utils.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	if raw {
		fmt.Println(out)
		return
	}
	lines := format(out)
	fmt.Println(strings.Join(lines, "\n\n"))
}
