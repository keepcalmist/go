package server

//Client is type with preassign time
type Client struct {
	time int
}

//ClientQueue is queue of clients to current column
type ClientQueue struct {
	next *ClientQueue
	prev *ClientQueue
	cl   Client
}

//Next returns Next client in queue
func (cl ClientQueue) Next() *ClientQueue {
	return cl.next
}

//Prev returns prev client in queue
func (cl ClientQueue) Prev() *ClientQueue {
	return cl.prev
}

//Len returns lenght of client queue
func (cl *ClientQueue) Len() int {
	var i int
	for cl.Next != nil {
		i++
	}
	return i
}
