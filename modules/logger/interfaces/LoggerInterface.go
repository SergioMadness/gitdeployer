package interfaces

type LoggerInterface interface {
	Log(category, message string)
	
	Flush()
	
	GetList() []LogRecordInterface
}