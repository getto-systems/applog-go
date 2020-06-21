package applog

import (
	"log"
	"os"
)

func Example_applog() {
	output := log.New(os.Stdout, "", 0)

	logger := NewDebugLogger(output)

	logger.Debug("debug: log message")
	// output: debug: log message
}
