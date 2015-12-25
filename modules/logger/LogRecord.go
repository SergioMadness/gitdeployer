package logger

type LogRecord struct {
	Time int64
	Hash string
}

// Constructor
func CreateLogRecord(time int64, hash string) *LogRecord {
	result := new(LogRecord)

	result.Hash = hash
	result.Time = time

	return result
}

func (lr *LogRecord) GetHash() string {
	return ""
}

func (lr *LogRecord) GetTimestamp() int64 {
	return 0
}
