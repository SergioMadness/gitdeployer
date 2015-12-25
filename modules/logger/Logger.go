package logger

import (
	"encoding/json"
	"gitdeployer/helpers"
	"io/ioutil"
	"os"
	"gitdeployer/modules/observer"
)

type Logger struct {
	observer.Observable
	messages []*LogMessage
}

// Constructor
func CreateLogger() *Logger {
	return new(Logger)
}

// Log event
func (l *Logger) Log(category, message string) {
	l.messages = append(l.messages, CreateLogMessage(category, message))
}

// Flush all messages to file
func (l *Logger) Flush() error {
	var tw []byte
	var err error

	fileName := helpers.RandomString(16)

	if !helpers.IsPathExists("logs") {
		os.MkdirAll("logs/", 0644)
	}

	if tw, err = json.Marshal(l.messages); err == nil {
		ioutil.WriteFile("logs/"+fileName+".data", tw, 0644)
	}

	return err
}

// Get log files
func (l *Logger) GetList() []LogRecord {
	var result []LogRecord

	return result
}
