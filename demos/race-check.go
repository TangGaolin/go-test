package main

import (
	"fmt"
	"strconv"
	"sync"
)

type MapData struct {
	Data map[string]string
	sync.Mutex
}

// 多个goroutine 使用race检测
func main() {
	//Demo1()
	//Demo2()
	//Demo3()
	//Demo4()
	correctDemo1()
}

func Demo1() {
	mapData := map[string]string{}
	mapData["data"] = "hello world"
	w := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			mapData["data"] = "hello world"
		}()
	}
	w.Wait()
}

func Demo2() {
	mapData := MapData{}
	mapData.Data = map[string]string{}
	w := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			mapData.Lock()
			mapData.Data["data"] = "hello world " + strconv.Itoa(i)
			fmt.Println(mapData.Data["data"])
			mapData.Unlock()
		}(i)
	}
	w.Wait()
}

func Demo3() {
	mapData := MapData{}
	mapData.Data = map[string]string{}
	w := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(mapData MapData) {
			defer w.Done()
			mapData.Lock()
			mapData.Data["data"] = "hello world"
			fmt.Println(mapData.Data["data"])
			mapData.Unlock()
		}(mapData)
	}
	w.Wait()
}

func Demo4() {
	type MapData2 struct {
		Data map[string]string
		sync.RWMutex
	}

	mapStr := map[string]string{}
	mapStr["key1"] = "value1"
	mapStr["key2"] = "value2"
	mapData := MapData2{}
	mapData.Data = map[string]string{}
	w := sync.WaitGroup{}

	for key, value := range mapStr {
		w.Add(1)
		go func(key string, value string) {
			defer w.Done()
			mapData.Lock()
			defer mapData.Unlock() //当len(mapStr) == 1的时候不加锁也可以通过
			mapData.Data[key] = value
		}(key, value)
	}
	w.Wait()
	fmt.Println(mapData)
}

func correctDemo1() {
	mapData := MapData{}
	mapData.Data = map[string]string{}
	w := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			mapData.Lock()
			mapData.Data["data"] = "hello world " + strconv.Itoa(i)
			fmt.Println(mapData.Data["data"])
			mapData.Unlock()
		}(i)
	}
	w.Wait()
}
