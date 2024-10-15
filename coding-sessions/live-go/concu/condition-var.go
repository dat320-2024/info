package concu

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// var mu sync.Mutex
var cond = sync.NewCond(&Mu)

//var wg sync.WaitGroup

const MAX_ITEMS = 10

var items = make([]int, 0)

func produce() {
	prod_count := 0
	for prod_count < 100 {
		item := rand.Intn(100)
		cond.L.Lock()
		if len(items) == MAX_ITEMS {
			fmt.Println("\t \t Buffer is full !!! , the producer is waiting")
			cond.Wait()
		}
		items = append(items, item)
		prod_count += 1
		fmt.Printf("Produced item id= %d, total items produced so far %d \n", item, prod_count)
		cond.Signal()
		cond.L.Unlock()
		time.Sleep(time.Millisecond * 10)
	}
	Wg.Done()
}
func consume() {
	consume_count := 0
	for consume_count < 100 {
		cond.L.Lock()
		if len(items) == 0 {
			fmt.Println("\t \t Buffer is empty !!, the consumer is waiting")
			cond.Wait()
		}
		item := items[0]
		items = items[1:]
		consume_count += 1
		fmt.Printf("Consumed item id= %d, total items consumed so far %d \n", item, consume_count)
		cond.Signal()
		cond.L.Unlock()

		time.Sleep(time.Millisecond * 500)
	}
	Wg.Done()
}
func RunCondVar() {
	Wg.Add(2)

	go produce()
	go consume()

	Wg.Wait()
}
