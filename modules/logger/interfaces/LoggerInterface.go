package interfaces

type LoggerInterface interface {
	Log(category, message string)
	
	Flush() error
}