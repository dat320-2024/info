package concu

import (
	"fmt"
	"time"
)

func RunPrinter() {
	S := func(a chan int, b chan chan int) { b <- a }
	P := func(a chan int) {
		fmt.Println("Printing", <-a)
	}
	C := func(b chan chan int) {
		a := <-b
		a <- 11
	}
	System := func() {
		a := make(chan int)
		b := make(chan chan int)
		go P(a)
		go C(b)
		go S(a, b)
	}
	System()
	time.Sleep(1 * time.Second)
}
