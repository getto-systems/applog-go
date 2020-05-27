package applog

// Logging AUDIT,ERROR level message
type ErrorLogger struct {
	config config
}

// New ErrorLogger
func NewErrorLogger(output Output, entry Entry) Logger {
	return ErrorLogger{
		config: config{
			output: output,
			entry:  entry,
		},
	}
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprint()
func (logger ErrorLogger) Audit(v ...interface{}) {
	logger.config.Audit(v...)
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprintf()
func (logger ErrorLogger) Auditf(format string, v ...interface{}) {
	logger.config.Auditf(format, v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprint()
func (logger ErrorLogger) Error(v ...interface{}) {
	logger.config.Error(v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprintf()
func (logger ErrorLogger) Errorf(format string, v ...interface{}) {
	logger.config.Errorf(format, v...)
}

// Drop messages
func (ErrorLogger) Warning(v ...interface{}) {
	// noop
}

// Drop messages
func (ErrorLogger) Warningf(format string, v ...interface{}) {
	// noop
}

// Drop messages
func (ErrorLogger) Info(v ...interface{}) {
	// noop
}

// Drop messages
func (ErrorLogger) Infof(format string, v ...interface{}) {
	// noop
}

// Drop messages
func (ErrorLogger) Debug(v ...interface{}) {
	// noop
}

// Drop messages
func (ErrorLogger) Debugf(format string, v ...interface{}) {
	// noop
}
