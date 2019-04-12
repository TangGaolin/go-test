// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

package main

import (
	"fmt"
	"time"
)

func main() {
	//demo1()

	demo2()
}

func demo1() {
	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func demo2() {
	chann := make(chan int)
	go func(chann chan int) {
		time.Sleep(1 * time.Second)
		chann <- 1
		close(chann)
	}(chann)

	for {
		select {
		case v, ok := <-chann:
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println("close")
				return
			}
		default:
			fmt.Println("waiting")
		}
	}
}
