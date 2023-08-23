package render

import (
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/log"
)

var r, _ = glamour.NewTermRenderer(
	// detect background color and pick either the default dark or light theme
	glamour.WithAutoStyle(),
	// wrap output at specific width (default is 80)
	glamour.WithWordWrap(40),
)

// Render markdown with gum
func Fancy(md string) {
	out, err := r.Render(md)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
