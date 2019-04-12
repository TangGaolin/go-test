package main

import (
	"fmt"
	"sync"
	"time"
)

//利用select 和 channel 实现队列管理

func main() {

	channs := make(chan int, 10)

	popDataFunc := func(i int) {
		select {
		case channs <- i:
			fmt.Println("data pop:", i)
		default:
			fmt.Println("data is full, pop failed:", i)
		}
	}

	//消费数据
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		waitTime := 10
		for {
			select {
			case msg := <-channs:
				fmt.Println("consumer:", msg)
				time.Sleep(2 * time.Second)
			default:
				fmt.Println("consumer waiting")
				waitTime--
			}

			if waitTime == 0 {
				break
			}
		}
		w.Done()
	}()

	//写入数据
	for i := 0; i < 30; i++ {
		fmt.Println("...........start pop:", i)
		popDataFunc(i)
		time.Sleep(1 * time.Second)
	}
	w.Wait()
}
