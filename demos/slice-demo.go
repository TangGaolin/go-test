package main

import "fmt"

func main() {
	sliceMen()
}

//slice 内存扩容规则1024内 是2倍扩容, 超过1024是1.25倍扩容
func sliceMen() {
	s5 := make([]int, 5)
	fmt.Println(len(s5), cap(s5))
	s5 = append(s5, 1)
	fmt.Println(len(s5), cap(s5))

	s6 := make([]int, 1024)
	fmt.Println(len(s6), cap(s6))
	s6 = append(s6, 1)
	fmt.Println(len(s6), cap(s6))
}
