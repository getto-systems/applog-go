package applog

import (
	"log"
	"os"
)

func ExampleLogger() {
	output := log.New(os.Stdout, "", 0)
	logger := NewDebugLogger(output, EntryTest)

	logger.Debug("log")
	// output: DEBUG: log
}
