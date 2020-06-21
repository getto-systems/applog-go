package applog

import (
	"log"
	"os"
)

func Example_debug() {
	output := log.New(os.Stdout, "", 0)
	logger := NewDebugLogger(output)

	logger.Always("always: log message")
	logger.Info("info: log message")
	logger.Debug("debug: log message")
	// output: always: log message
	// info: log message
	// debug: log message
}

func Example_info() {
	output := log.New(os.Stdout, "", 0)
	logger := NewInfoLogger(output)

	logger.Always("always: log message")
	logger.Info("info: log message")
	logger.Debug("debug: log message")
	// output: always: log message
	// info: log message
}

func Example_quiet() {
	output := log.New(os.Stdout, "", 0)
	logger := NewQuietLogger(output)

	logger.Always("always: log message")
	logger.Info("info: log message")
	logger.Debug("debug: log message")
	// output: always: log message
}
