package observer

import (
	"gitdeployer/modules/observer/interfaces"
)

type Observable struct {
	observers []*Observer
}

func (o *Observable) AddObserver(obs *Observer) {
	o.observers = append(o.observers, obs)
}

func (o *Observable) GetObservers() []*Observer {
	return o.observers
}

func (o *Observable) RemoveObserver(obs *interfaces.Observer) {

}

func (o *Observable) Invoke(message interface{}) {
	for _, obs := range o.observers {
		obs.Notify(message)
	}
}
