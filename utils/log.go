package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var _ERROR = color.RedString("ERROR:")

// log an error message
func ErrorF(err error) {
	fmt.Fprintln(os.Stderr, _ERROR, err)
}
