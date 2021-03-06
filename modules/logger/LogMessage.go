package logger

import (
	"time"
)

type LogMessage struct {
	Time     int64
	Category string
	Message  string
}

// Constructor
func CreateLogMessage(category, message string) *LogMessage {
	result := new(LogMessage)

	result.Time = time.Now().Unix()
	result.Category = category
	result.Message = message

	return result
}
