package observer

import(
	"gitdeployer/modules/observer/interfaces"
)

type Observable struct {
	observers []*interfaces.Observer
}

func (o *Observable) AddObserver(obs *interfaces.Observer) {
	o.observers = append(o.observers, obs)
}

func (o *Observable) RemoveObserver(obs *interfaces.Observer) {
	
}