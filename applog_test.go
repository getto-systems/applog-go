package applog

import "testing"

import (
	"fmt"
	"reflect"
)

type OutputTest struct {
	messages []string
}

func (output *OutputTest) Println(v ...interface{}) {
	output.messages = append(output.messages, fmt.Sprint(v...))
}

func EntryTest(level string, message string) string {
	return fmt.Sprintf("%s: %s", level, message)
}

func TestDebugLogger(t *testing.T) {
	output := &OutputTest{}
	logger := NewDebugLogger(output, EntryTest)

	logger.Audit("log")
	logger.Auditf("log: %s", "message")

	logger.Error("log")
	logger.Errorf("log: %s", "message")

	logger.Warning("log")
	logger.Warningf("log: %s", "message")

	logger.Info("log")
	logger.Infof("log: %s", "message")

	logger.Debug("log")
	logger.Debugf("log: %s", "message")

	expected := []string{
		"AUDIT: log",
		"AUDIT: log: message",

		"ERROR: log",
		"ERROR: log: message",

		"WARNING: log",
		"WARNING: log: message",

		"INFO: log",
		"INFO: log: message",

		"DEBUG: log",
		"DEBUG: log: message",
	}

	if !reflect.DeepEqual(output.messages, expected) {
		t.Errorf("output did not expected!\noutput:   %v\nexpected: %v", output.messages, expected)
	}
}

func TestInfoLogger(t *testing.T) {
	output := &OutputTest{}
	logger := NewInfoLogger(output, EntryTest)

	logger.Audit("log")
	logger.Auditf("log: %s", "message")

	logger.Error("log")
	logger.Errorf("log: %s", "message")

	logger.Warning("log")
	logger.Warningf("log: %s", "message")

	logger.Info("log")
	logger.Infof("log: %s", "message")

	logger.Debug("log")
	logger.Debugf("log: %s", "message")

	expected := []string{
		"AUDIT: log",
		"AUDIT: log: message",

		"ERROR: log",
		"ERROR: log: message",

		"WARNING: log",
		"WARNING: log: message",

		"INFO: log",
		"INFO: log: message",

		//"DEBUG: log",
		//"DEBUG: log: message",
	}

	if !reflect.DeepEqual(output.messages, expected) {
		t.Errorf("output did not expected!\noutput:   %v\nexpected: %v", output.messages, expected)
	}
}

func TestWarningLogger(t *testing.T) {
	output := &OutputTest{}
	logger := NewWarningLogger(output, EntryTest)

	logger.Audit("log")
	logger.Auditf("log: %s", "message")

	logger.Error("log")
	logger.Errorf("log: %s", "message")

	logger.Warning("log")
	logger.Warningf("log: %s", "message")

	logger.Info("log")
	logger.Infof("log: %s", "message")

	logger.Debug("log")
	logger.Debugf("log: %s", "message")

	expected := []string{
		"AUDIT: log",
		"AUDIT: log: message",

		"ERROR: log",
		"ERROR: log: message",

		"WARNING: log",
		"WARNING: log: message",

		//"INFO: log",
		//"INFO: log: message",

		//"DEBUG: log",
		//"DEBUG: log: message",
	}

	if !reflect.DeepEqual(output.messages, expected) {
		t.Errorf("output did not expected!\noutput:   %v\nexpected: %v", output.messages, expected)
	}
}

func TestErrorLogger(t *testing.T) {
	output := &OutputTest{}
	logger := NewErrorLogger(output, EntryTest)

	logger.Audit("log")
	logger.Auditf("log: %s", "message")

	logger.Error("log")
	logger.Errorf("log: %s", "message")

	logger.Warning("log")
	logger.Warningf("log: %s", "message")

	logger.Info("log")
	logger.Infof("log: %s", "message")

	logger.Debug("log")
	logger.Debugf("log: %s", "message")

	expected := []string{
		"AUDIT: log",
		"AUDIT: log: message",

		"ERROR: log",
		"ERROR: log: message",

		//"WARNING: log",
		//"WARNING: log: message",

		//"INFO: log",
		//"INFO: log: message",

		//"DEBUG: log",
		//"DEBUG: log: message",
	}

	if !reflect.DeepEqual(output.messages, expected) {
		t.Errorf("output did not expected!\noutput:   %v\nexpected: %v", output.messages, expected)
	}
}
