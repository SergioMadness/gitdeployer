package interfaces

type Observer interface {
	Notify(message interface{})
}