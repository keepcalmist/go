package server

import (
	"fmt"
	"time"
)

type GasStation struct {
	ip                 string
	port               string
	gasColumnList      []*GasColumn
	queueForEachColumn map[GasColumn]*ClientQueue
}

type GasColumn struct {
	number   int
	workTime int
	maxQueue int
	client   Client
	enable   bool
}

func LoadGasStation(ip string, port string) (GS *GasStation) {
	GS = &GasStation{
		ip:                 ip,
		port:               port,
		gasColumnList:      make([]*GasColumn, 1),
		queueForEachColumn: make(map[GasColumn]*ClientQueue),
	}
	return GS
}

type SettingsGS interface {
	EmptyGS() (GS *GasStation)
	ChangeIP(string)
	ChagePort(string)
}

func EmptyGS() *GasStation {
	return &GasStation{
		ip:                 "",
		port:               "",
		gasColumnList:      make([]*GasColumn, 1),
		queueForEachColumn: make(map[GasColumn]*ClientQueue),
	}
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
		fmt.Printf("%-3d|", gs.queueForEachColumn[*i].Len())
	}
}

//AddColumn is func to add new column to gasColumnList in gasStation
func (gs *GasStation) AddColumn() {
	newGC := &GasColumn{
		number:   len(gs.gasColumnList),
		workTime: 0,
		maxQueue: 1,
		enable:   false,
	}
	gs.gasColumnList = append(gs.gasColumnList, newGC)
	gs.addColumnToQueue(*newGC)
}

//ChangeAble is func to swap able to other(if gascolumn have true able then it will has false and conversely)
func (gc *GasColumn) ChangeAble() {
	if gc.enable == true {
		gc.enable = false
	} else if gc.enable == false {
		gc.enable = true
	}
	fmt.Printf("Able has been changed successfully to %b", gc.enable)
}

func (gs *GasStation) addColumnToQueue(gc GasColumn) {
	var cl *ClientQueue
	gs.queueForEachColumn[gc] = append(gs.queueForEachColumn[gc], cl)
}

//ChangeMaxQueueColumn is func to change gascolumn max queue
func (gs *GasStation) ChangeMaxQueueColumn(n int, max int) {
	gs.gasColumnList[n-1].maxQueue = max
}

func (gs GasStation) countLenQueue(gc GasColumn) int {
	return gs.queueForEachColumn[gc].Len()
}

func (gs GasStation) countWT(num GasColumn) (sum int, err error) {
	queue := gs.queueForEachColumn[num]
	for queue.Next() != nil {
		sum += queue.cl.time
	}
	return sum, nil
}

func (gs *GasStation) ChangeGCWT(gc *GasColumn) (err error) {
	gc.workTime, err = gs.countWT(*gc)
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

func Start() {
	lo
}
