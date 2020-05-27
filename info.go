package applog

// Logging AUDIT,ERROR,WARNING,INFO level message
type InfoLogger struct {
	config config
}

// New InfoLogger
func NewInfoLogger(output Output, entry Entry) Logger {
	return InfoLogger{
		config: config{
			output: output,
			entry:  entry,
		},
	}
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprint()
func (logger InfoLogger) Audit(v ...interface{}) {
	logger.config.Audit(v...)
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprintf()
func (logger InfoLogger) Auditf(format string, v ...interface{}) {
	logger.config.Auditf(format, v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprint()
func (logger InfoLogger) Error(v ...interface{}) {
	logger.config.Error(v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprintf()
func (logger InfoLogger) Errorf(format string, v ...interface{}) {
	logger.config.Errorf(format, v...)
}

// Logging WARNING level message.
// Argument pass through fmt.Sprint()
func (logger InfoLogger) Warning(v ...interface{}) {
	logger.config.Warning(v...)
}

// Logging WARNING level message.
// Argument pass through fmt.Sprintf()
func (logger InfoLogger) Warningf(format string, v ...interface{}) {
	logger.config.Warningf(format, v...)
}

// Logging INFO level message.
// Argument pass through fmt.Sprint()
func (logger InfoLogger) Info(v ...interface{}) {
	logger.config.Info(v...)
}

// Logging INFO level message.
// Argument pass through fmt.Sprintf()
func (logger InfoLogger) Infof(format string, v ...interface{}) {
	logger.config.Infof(format, v...)
}

// Drop messages
func (InfoLogger) Debug(v ...interface{}) {
	// noop
}

// Drop messages
func (InfoLogger) Debugf(format string, v ...interface{}) {
	// noop
}
