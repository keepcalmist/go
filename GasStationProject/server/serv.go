package server

import (
	"fmt"
)

//GasStation is main object in the project
type GasStation struct {
	ip                 string
	port               string
	gasColumnList      []*GasColumn
	queueForEachColumn map[GasColumn]*ClientQueue
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
func (gs *GasStation) addColumnToQueue(gc GasColumn) {
	var cl *ClientQueue
	gs.queueForEachColumn[gc] = cl
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

func Start() {

}
