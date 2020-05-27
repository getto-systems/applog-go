package applog

// Logging all level message
type DebugLogger struct {
	config config
}

// New DebugLogger
func NewDebugLogger(output Output, entry Entry) Logger {
	return DebugLogger{
		config: config{
			output: output,
			entry:  entry,
		},
	}
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprint()
func (logger DebugLogger) Audit(v ...interface{}) {
	logger.config.Audit(v...)
}

// Logging AUDIT level message.
// Argument pass through fmt.Sprintf()
func (logger DebugLogger) Auditf(format string, v ...interface{}) {
	logger.config.Auditf(format, v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprint()
func (logger DebugLogger) Error(v ...interface{}) {
	logger.config.Error(v...)
}

// Logging ERROR level message.
// Argument pass through fmt.Sprintf()
func (logger DebugLogger) Errorf(format string, v ...interface{}) {
	logger.config.Errorf(format, v...)
}

// Logging WARNING level message.
// Argument pass through fmt.Sprint()
func (logger DebugLogger) Warning(v ...interface{}) {
	logger.config.Warning(v...)
}

// Logging WARNING level message.
// Argument pass through fmt.Sprintf()
func (logger DebugLogger) Warningf(format string, v ...interface{}) {
	logger.config.Warningf(format, v...)
}

// Logging INFO level message.
// Argument pass through fmt.Sprint()
func (logger DebugLogger) Info(v ...interface{}) {
	logger.config.Info(v...)
}

// Logging INFO level message.
// Argument pass through fmt.Sprintf()
func (logger DebugLogger) Infof(format string, v ...interface{}) {
	logger.config.Infof(format, v...)
}

// Logging DEBUG level message.
// Argument pass through fmt.Sprint()
func (logger DebugLogger) Debug(v ...interface{}) {
	logger.config.Debug(v...)
}

// Logging DEBUG level message.
// Argument pass through fmt.Sprintf()
func (logger DebugLogger) Debugf(format string, v ...interface{}) {
	logger.config.Debugf(format, v...)
}
