package applog

// Logging only always level message
type QuietLogger struct {
	config config
}

// New QuietLogger
func NewQuietLogger(output Output) Logger {
	return QuietLogger{
		config: config{
			output: output,
		},
	}
}

// Logging message.
func (logger QuietLogger) Always(v interface{}) {
	logger.config.Println(v)
}

// Drop info level message.
func (QuietLogger) Info(v interface{}) {
	// noop
}

// Drop debug level message.
func (QuietLogger) Debug(v interface{}) {
	// noop
}
