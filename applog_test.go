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

func TestDebugLogger(t *testing.T) {
	output := &OutputTest{}
	logger := NewDebugLogger(output)

	logger.Always("always: log")
	logger.Info("info: log")
	logger.Debug("debug: log")

	expected := []string{
		"always: log",
		"info: log",
		"debug: log",
	}

	if !reflect.DeepEqual(output.messages, expected) {
		t.Errorf("output did not expected!\noutput:   %v\nexpected: %v", output.messages, expected)
	}
}

func TestInfoLogger(t *testing.T) {
	output := &OutputTest{}
	logger := NewInfoLogger(output)

	logger.Always("always: log")
	logger.Info("info: log")
	logger.Debug("debug: log")

	expected := []string{
		"always: log",
		"info: log",
		//"DEBUG: log",
	}

	if !reflect.DeepEqual(output.messages, expected) {
		t.Errorf("output did not expected!\noutput:   %v\nexpected: %v", output.messages, expected)
	}
}

func TestQuietLogger(t *testing.T) {
	output := &OutputTest{}
	logger := NewQuietLogger(output)

	logger.Always("always: log")
	logger.Info("info: log")
	logger.Debug("debug: log")

	expected := []string{
		"always: log",
		//"info: log",
		//"debug: log",
	}

	if !reflect.DeepEqual(output.messages, expected) {
		t.Errorf("output did not expected!\noutput:   %v\nexpected: %v", output.messages, expected)
	}
}
