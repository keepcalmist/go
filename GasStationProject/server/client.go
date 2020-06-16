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
func (clQ ClientQueue) Next() *ClientQueue {
	return clQ.next
}

//Prev returns prev client in queue
func (clQ ClientQueue) Prev() *ClientQueue {
	return clQ.prev
}

//Len returns lenght of client queue
func (clQ *ClientQueue) Len() int {
	var i int
	for clQ.cl.time != 0 {
		i++
		clQ.Next()
	}
	return i
}

//SetRefillTime if func to set client's refilling time
func (clQ *ClientQueue) SetRefillTime(time int) {
	clQ.cl.time = time
}

func CreateRootCl() (clQ *ClientQueue) {
	return &ClientQueue{
		next: nil,
		prev: nil,
		cl:   Client{time: 0},
	}
}

func (clQ *ClientQueue) CreateQueue(value int) {
	for i := 0; i < value; i++ {
		clQ.next = CreateRootCl()
		clQ.next.prev = clQ
		clQ.Next()
	}
}
