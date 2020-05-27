package applog

// Logging AUDIT,ERROR,WARNING level message
type WarnLogger struct {
	config config
}

// New WarnLogger
func NewWarnLogger(output Output, entry Entry) Logger {
	return WarnLogger{
		config: config{
			output: output,
			entry:  entry,
		},
	}
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprint()
func (logger WarnLogger) Audit(v ...interface{}) {
	logger.config.Audit(v...)
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprintf()
func (logger WarnLogger) Auditf(format string, v ...interface{}) {
	logger.config.Auditf(format, v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprint()
func (logger WarnLogger) Error(v ...interface{}) {
	logger.config.Error(v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprintf()
func (logger WarnLogger) Errorf(format string, v ...interface{}) {
	logger.config.Errorf(format, v...)
}

// Logging WARNING level message.
// Argument pass through fmt.Sprint()
func (logger WarnLogger) Warn(v ...interface{}) {
	logger.config.Warn(v...)
}

// Logging WARNING level message.
// Argument pass through fmt.Sprintf()
func (logger WarnLogger) Warnf(format string, v ...interface{}) {
	logger.config.Warnf(format, v...)
}

// Drop messages
func (WarnLogger) Info(v ...interface{}) {
	// noop
}

// Drop messages
func (WarnLogger) Infof(format string, v ...interface{}) {
	// noop
}

// Drop messages
func (WarnLogger) Debug(v ...interface{}) {
	// noop
}

// Drop messages
func (WarnLogger) Debugf(format string, v ...interface{}) {
	// noop
}
