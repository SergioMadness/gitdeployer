package logger

import (
	"time"
)

type LogMessage struct {
	Time     int64
	Category string
	Message  string
}

func CreateLogMessage(category, message string) *LogMessage {
	result := new(LogMessage)

	result.Time = time.Now().Unix()
	result.Category = category
	result.Message = category

	return result
}
