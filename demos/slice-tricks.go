package main

import (
	"fmt"
	"math/rand"
)

/*
一些slice的使用技巧 from:https://github.com/golang/go/wiki/SliceTricks
*/
func main() {
	appendVector()
	cut()
	insert()
	shift()
	reverse()
	shuffle()
}

//两个数组合并
func appendVector() {
	a := []int{1, 2}
	b := []int{3, 4}
	a = append(a, b...)
	fmt.Println("appendVector", a)
}

//减掉中间的元素
func cut() {
	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a[:3], a[5:]...)
	fmt.Println("cut", a)
}

//指定位置插入
func insert() {
	a := []int{1, 2, 3, 4, 6}
	insertPos := 4
	fmt.Println("insert before:", a)
	a = append(a[:insertPos], append([]int{5}, a[insertPos:]...)...)
	fmt.Println("insert after:", a)
}

//去除数组中的第一个数
func shift() {
	var x int
	a := []int{1, 2, 3, 4}
	x, a = a[0], a[1:]
	fmt.Println("pop data:", x)
	fmt.Println("res data:", a)
}

//数组反转
func reverse() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("reverse before", a)
	for i := 0; i <= len(a)/2-1; i++ {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	fmt.Println("reverse after", a)

	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	fmt.Println("reverse after again", a)
}

//随机打散顺序
func shuffle() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("shuffle before", a)
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println("shuffle after", a)
}
