package applog

// Logging except debug level message
type DebugLogger struct {
	config config
}

// New DebugLogger
func NewDebugLogger(output Output) Logger {
	return DebugLogger{
		config: config{
			output: output,
		},
	}
}

// Logging message.
func (logger DebugLogger) Always(v interface{}) {
	logger.config.Println(v)
}

// Logging message.
func (logger DebugLogger) Info(v interface{}) {
	logger.config.Println(v)
}

// Logging message.
func (logger DebugLogger) Debug(v interface{}) {
	logger.config.Println(v)
}
