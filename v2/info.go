package applog

// Logging except debug level message
type InfoLogger struct {
	config config
}

// New InfoLogger
func NewInfoLogger(output Output) Logger {
	return InfoLogger{
		config: config{
			output: output,
		},
	}
}

// Logging message.
func (logger InfoLogger) Always(v interface{}) {
	logger.config.Println(v)
}

// Logging message.
func (logger InfoLogger) Info(v interface{}) {
	logger.config.Println(v)
}

// Drop debug level message.
func (InfoLogger) Debug(v interface{}) {
	// noop
}
