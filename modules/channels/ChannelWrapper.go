package channels

import "gitdeployer/modules/observer/interfaces"

type ChannelWrapper struct {
	interfaces.Observer
	channels map[int]chan string
}

var currentWrapper *ChannelWrapper

func GetWrapper() *ChannelWrapper {
	if currentWrapper == nil {
		currentWrapper = new(ChannelWrapper)
		currentWrapper.channels = make(map[int]chan string)
	}

	return currentWrapper
}

func (cw *ChannelWrapper) GetChannel(id int) chan string {
	if cw.channels[id] == nil {
		cw.channels[id] = make(chan string)
	}
	return cw.channels[id]
}

func (cw *ChannelWrapper) Notify(message interface{}) {
	for _, channel := range cw.channels {
		channel <- message.(string)
	}
}
