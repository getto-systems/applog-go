package applog

import (
	"fmt"
	"log"
	"os"
)

func Example_applog() {
	output := log.New(os.Stdout, "", 0)
	entry := func(level string, message string) string {
		return fmt.Sprintf("%s: %s", level, message)
	}

	logger := NewDebugLogger(output, entry)

	logger.Debug("log message")
	// output: DEBUG: log message
}
