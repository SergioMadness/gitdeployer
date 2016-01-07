package interfaces

type Observable interface {
	AddObserver(o *Observer)
	RemoveObserver(o *Observer)
	Invoke(message interface{})
}