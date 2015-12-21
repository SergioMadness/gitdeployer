package interfaces

type LogRecordInterface interface {
	GetHash() string
	
	GetTimestamp() int64
}