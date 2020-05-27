package applog

import (
	"fmt"
)

// Logging message
type Logger interface {
	Audit(v ...interface{})
	Auditf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
}

// Logging to real world
type Output interface {
	Println(v ...interface{})
}

// Format log-entry from level-string and message
type Entry func(level string, message string) string

const (
	level_audit   = "AUDIT"
	level_error   = "ERROR"
	level_warning = "WARNING"
	level_info    = "INFO"
	level_debug   = "DEBUG"
)

type config struct {
	output Output
	entry  Entry
}

func (config config) Audit(v ...interface{}) {
	config.Message(level_audit, v...)
}
func (config config) Auditf(format string, v ...interface{}) {
	config.Messagef(level_audit, format, v...)
}

func (config config) Debug(v ...interface{}) {
	config.Message(level_debug, v...)
}
func (config config) Debugf(format string, v ...interface{}) {
	config.Messagef(level_debug, format, v...)
}

func (config config) Info(v ...interface{}) {
	config.Message(level_info, v...)
}
func (config config) Infof(format string, v ...interface{}) {
	config.Messagef(level_info, format, v...)
}

func (config config) Warn(v ...interface{}) {
	config.Message(level_warning, v...)
}
func (config config) Warnf(format string, v ...interface{}) {
	config.Messagef(level_warning, format, v...)
}

func (config config) Error(v ...interface{}) {
	config.Message(level_error, v...)
}
func (config config) Errorf(format string, v ...interface{}) {
	config.Messagef(level_error, format, v...)
}

func (config config) Message(level string, v ...interface{}) {
	config.Println(level, fmt.Sprint(v...))
}
func (config config) Messagef(level string, format string, v ...interface{}) {
	config.Println(level, fmt.Sprintf(format, v...))
}
func (config config) Println(level string, message string) {
	config.output.Println(config.entry(level, message))
}
