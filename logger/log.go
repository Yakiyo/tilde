package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

// initialization of logger
func Init() {
	// set log level to warn
	log.SetLevel(log.WarnLevel)

	log.SetReportTimestamp(false)
	log.SetOutput(os.Stderr)
}
