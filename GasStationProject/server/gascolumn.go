package server

import (
	"fmt"
	"time"
)

type GasColumn struct {
	number     int
	workTime   int
	maxQueue   int
	client     *ClientQueue
	lastClient *ClientQueue
	enable     bool
}

func (gc *GasColumn) refill() error {
	ticker := time.NewTicker(time.Duration(gc.client.cl.time) * time.Second)
	done := make(chan bool)
	go func(n int) {
		i := 0
		for {
			select {
			case <-done:
				fmt.Printf("GasColumn number: %d;\nProccesing refilling %%%d ", gc.number, 100)
				return
			case <-ticker.C:
				i++
				fmt.Printf("GasColumn number: %d;\nProccesing refilling %%%d ", gc.number, i/n)
			}
		}
	}(gc.client.cl.time / 100)
	time.Sleep((time.Duration(gc.client.cl.time)) * time.Second)
	ticker.Stop()
	done <- true
	return nil
}

//RefilProc is func to refill car and delete them from queue
func (gc *GasColumn) RefilProc() error {
	err := gc.refill()
	if err != nil {
		return err
	}
	gc.client = gc.client.Next()
	gc.client.prev.next = nil
	gc.client.prev = nil
	return nil
}

func (gc *GasColumn) AddClient(cl *Client) {
	gc.lastClient.next = &ClientQueue{
		next: nil,
		prev: gc.lastClient,
		cl:   *cl,
	}
	gc.lastClient = gc.lastClient.Next()
}

func (gc *GasColumn) ChangeAble() {
	if gc.enable == true {
		gc.enable = false
	} else if gc.enable == false {
		gc.enable = true
	}
	fmt.Printf("Able has been changed successfully to %b", gc.enable)
}
