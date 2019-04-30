package main

import (
	"fmt"
)

/******
什么是闭包：
是由函数及其相关引用环境组合而成的实体,即：闭包=函数+引用环境
可以直接调用或者赋值于某个变量的匿名函数，称为闭包。
Golang中，所有的匿名函数都是的闭包。

闭包引用了x变量，acc1,acc2可看作2个不同的实例，实例之间互不影响。实例内部，x变量是同一个地址，因此具有“累加效应”。
******/
func main() {
	var acc1 = accumulator()
	acc1(1)
	acc1(10)
	acc1(100)

	var acc2 = accumulator()
	acc2(2)
	acc2(20)
	acc2(200)
}

func accumulator() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Printf("(x addr:%v, x value:%v) ", &x, x)
		x += delta
		fmt.Printf("- %d\n", x)
		return x
	}
}