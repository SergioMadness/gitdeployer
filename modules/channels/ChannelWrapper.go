package channels

var channels = make(map[int]chan string)

func GetChannel(id int) chan string {
	if channels[id]==nil {
		channels[id] = make(chan string)
	}
	return channels[id]
}