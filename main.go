// Fast tldr console client
package main

import (
	"github.com/Yakiyo/tilde/cmd"
	"github.com/charmbracelet/log"
)

func main() {
	defer panicHandler()

	cmd.Execute()
}

// handle unexpected panics
func panicHandler() {
	if err := recover(); err != nil {
		log.Error("Unhandled error cause panic. Please consider filing a bug at https://github.com/Yakiyo/tilde")
		log.Fatal(err)
	}
}