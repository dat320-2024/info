package concu

import (
	"fmt"
	"sync"
	_ "time"
)

var Mu sync.Mutex
var Wg sync.WaitGroup

func count(c *int, ite int, safe bool) {
	for i := 0; i < ite; i++ {
		if safe {
			Mu.Lock()
			//heavy increment NOT atomic
			//*c *= 2
			//time.Sleep(time.Microsecond)
			//*c /= 2
			*c++
			Mu.Unlock()
		} else {
			//heavy increment NOT atomic
			//*c *= 2
			//time.Sleep(time.Microsecond)
			//*c /= 2
			*c++
		}

	}
	Wg.Done()

}

func RunRC() {
	iterations := 5
	number_of_threads := 500
	expected_count := iterations * number_of_threads
	fmt.Println("Expected count =", expected_count)

	Wg.Add(number_of_threads)

	counter := 0
	for i := 0; i < number_of_threads; i++ {
		go count(&counter, iterations, false)
	}
	Wg.Wait()
	fmt.Println("Unsafe count =", counter, " Race condition !")

	Wg.Add(number_of_threads)
	counter = 0
	for i := 0; i < number_of_threads; i++ {
		go count(&counter, iterations, false)
	}

	Wg.Wait()

	fmt.Println("Safe count =", counter)
}
