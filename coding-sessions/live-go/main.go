package main

import (
	"fmt"

	//"github.com/dat320-2024/info/coding-sessions/live-go/caching"

	"github.com/dat320-2024/info/coding-sessions/live-go/concu"
	"github.com/dat320-2024/info/coding-sessions/live-go/scheduling"
)

func main() {
	//fmt.Println("DAT320 terminal")
	//utils.DemoQueue()
	//utils.SliceFromArrayDemo()
	//utils.SliceMakeDemo()
	//terminal := terminal.NewTM()
	//name := os.Args[1]
	//terminal.Run(name)

	//RunLottery()
	//m := caching.Create_Matrix(1000, 4096)
	//caching.LoopByRows(*m)
	//caching.LoopByColumns(*m)

	//con.RunRC()
	//concu.RunAgentDemo()
	//concu.RunPrinter()
	concu.RunCondVar()
}

func RunLottery() {

	slots := 10
	fmt.Print("Lottery scheduling Demo\n\n")

	l1 := scheduling.NewLottery(1, 10)
	l2 := scheduling.NewLottery(2, 10)
	l3 := scheduling.NewLottery(3, 10)
	l4 := scheduling.NewLottery(4, 10)
	l5 := scheduling.NewLottery(5, 10)

	scheduleMap := scheduling.Schedule(slots, []*scheduling.Lottery{l1, l2, l3, l4, l5})

	fmt.Println("Schedule Map:")
	fmt.Printf("%-10s %-10s %-10s\n", "PID", "Count", " %")
	for key, value := range scheduleMap {
		fmt.Printf("%-10d %-10d %-10d\n", key, value, value*100/slots)
	}
}
