package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {
		once.Do(onced)
	}
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
