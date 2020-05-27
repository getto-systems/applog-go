package applog

// Logging AUDIT,ERROR,WARNING level message
type WarningLogger struct {
	config config
}

// New WarningLogger
func NewWarningLogger(output Output, entry Entry) Logger {
	return WarningLogger{
		config: config{
			output: output,
			entry:  entry,
		},
	}
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprint()
func (logger WarningLogger) Audit(v ...interface{}) {
	logger.config.Audit(v...)
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprintf()
func (logger WarningLogger) Auditf(format string, v ...interface{}) {
	logger.config.Auditf(format, v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprint()
func (logger WarningLogger) Error(v ...interface{}) {
	logger.config.Error(v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprintf()
func (logger WarningLogger) Errorf(format string, v ...interface{}) {
	logger.config.Errorf(format, v...)
}

// Logging WARNING level message.
// Argument pass through fmt.Sprint()
func (logger WarningLogger) Warning(v ...interface{}) {
	logger.config.Warning(v...)
}

// Logging WARNING level message.
// Argument pass through fmt.Sprintf()
func (logger WarningLogger) Warningf(format string, v ...interface{}) {
	logger.config.Warningf(format, v...)
}

// Drop messages
func (WarningLogger) Info(v ...interface{}) {
	// noop
}

// Drop messages
func (WarningLogger) Infof(format string, v ...interface{}) {
	// noop
}

// Drop messages
func (WarningLogger) Debug(v ...interface{}) {
	// noop
}

// Drop messages
func (WarningLogger) Debugf(format string, v ...interface{}) {
	// noop
}
