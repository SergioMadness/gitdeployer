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

func CreateLogger() *Logger {
	return new(Logger)
}

func (l *Logger) Log(category, message string) {
	l.messages = append(l.messages, CreateLogMessage(category, message))
}

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

func (l *Logger) GetList() []LogRecord {
	var result []LogRecord

	return result
}
