package server

import (
	"fmt"
	"sync"
	"time"
)

type Client struct {
	//mu   sync.Mutex
	time int
}

type GasStation struct {
	ip                 string
	port               string
	gasColumnList      []*GasColumn
	queueForEachColumn map[GasColumn][]*Client
}

type GasColumn struct {
	number   int
	workTime int
	maxQueue int
	mu       sync.Mutex
	client   Client
}

func LoadGasStation(ip string, port string) (GS *GasStation) {
	GS = &GasStation{
		ip:                 ip,
		port:               port,
		gasColumnList:      make([]*GasColumn, 1),
		queueForEachColumn: make(map[GasColumn][]*Client),
	}
	return GS
}

type SettingsGS interface {
	EmptyGS() (GS *GasStation)
	ChangeIP(string)
	ChagePort(string)
}

func EmptyGS() (GS *GasStation) {
	return new(GasStation)
}

func (gs *GasStation) ChangeIP(ip string) {
	gs.ip = ip
}

func (gs *GasStation) ChangePort(port string) {
	gs.port = port
}

func (gs *GasStation) PrintColumns() {
	fmt.Printf("       |")
	for _, i := range gs.gasColumnList {
		fmt.Printf("%-3d|", i)
	}

	fmt.Printf("\nQueue: |")
	for _, i := range gs.gasColumnList {
		fmt.Printf("%-3d|", len(gs.queueForEachColumn[*i]))
	}
}

//AddColumn is func to add new column to gasColumnList in gasStation
func (gs *GasStation) AddColumn() {
	newGC := &GasColumn{
		number:   len(gs.gasColumnList),
		workTime: 0,
		maxQueue: 1,
	}
	gs.gasColumnList = append(gs.gasColumnList, newGC)
	gs.addColumnToQueue(*newGC)
}

func (gs *GasStation) addColumnToQueue(gc GasColumn) {
	var cl *Client
	gs.queueForEachColumn[gc] = append(gs.queueForEachColumn[gc], cl)
}

//ChangeQueueColumn is func to change gascolumn max queue
func (gs *GasStation) ChangeMaxQueueColumn(n int, max int) {
	gs.gasColumnList[n-1].maxQueue = max
}

func (gs GasStation) countLenQueue(gc GasColumn) int {
	return len(gs.queueForEachColumn[gc])
}

func (gs GasStation) countWT(num GasColumn) (sum int, err error) {
	queue := gs.queueForEachColumn[num]
	for _, i := range queue {
		sum += i.time
	}
	return sum, nil
}

func (gs *GasStation) ChangeGCWT(gc *GasColumn) (err error) {
	gc.mu.Lock()
	gc.workTime, err = gs.countWT(*gc)
	gc.mu.Unlock()
	return err
}

func (gc *GasColumn) refill() error {
	ticker := time.NewTicker(time.Duration(gc.client.time) * time.Second)
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
	}(gc.client.time / 100)
	time.Sleep((time.Duration(gc.client.time)) * time.Second)
	ticker.Stop()
	done <- true
	return nil
}
