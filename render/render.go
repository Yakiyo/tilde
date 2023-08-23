package render

import (
	"fmt"

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

}
