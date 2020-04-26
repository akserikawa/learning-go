package main

import (
	"fmt"
	"time"
)

func genNumbers(from, to int, ch chan int) {
	for j := from; j <= to; j++ {
		ch <- j
		fmt.Printf("Published: %d\n", j)
	}
}

func timeout(n int, ch chan bool) {
	time.Sleep(time.Duration(n) * time.Second)
	ch <- true
}

func main() {
	from, to := 1, 10
	genCh := make(chan int)
	timeoutCh := make(chan bool)

	go genNumbers(from, to, genCh)
	go timeout(1, timeoutCh)

	for {
		select {
		case x := <-genCh:
			fmt.Printf("Received: %d\n", x)
		case <-timeoutCh:
			fmt.Println("timeout")
			return
		}
	}
}
