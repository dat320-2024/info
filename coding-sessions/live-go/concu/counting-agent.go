package concu

import (
	"fmt"
	"math/rand"
	"time"
)

type SearchingAgent struct {
	Link       chan int
	Multiple   int
	Iterations int
	Max        int
	Name       string
	Exit       chan int
}

func (ag *SearchingAgent) Search() {
	v := 0
	for i := 0; i < ag.Iterations; i++ {
		v = rand.Intn(ag.Max)
		if v%ag.Multiple == 0 {
			//found it
			fmt.Println(ag.Name, "found", v, "Winner")
			ag.Link <- v
			ag.Exit <- 1 //some value
			return

		}
		select {
		case <-ag.Link:
			fmt.Println(ag.Name, "Lost")
			ag.Exit <- 1 //some value
			return
		default:
			time.Sleep(time.Millisecond)
		}

	}
	ag.Link <- 1
	ag.Exit <- 1
}
func RunAgentDemo() {
	link := make(chan int)
	exit := make(chan int)
	agent1 := SearchingAgent{Name: "A1", Multiple: 110, Max: 3000, Iterations: 300, Link: link, Exit: exit}
	agent2 := SearchingAgent{Name: "A2", Multiple: 110, Max: 3000, Iterations: 300, Link: link, Exit: exit}
	go agent1.Search()
	go agent2.Search()
	<-exit
	<-exit
	fmt.Println("Done playing")

}
