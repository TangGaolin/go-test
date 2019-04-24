package main

import "fmt"

func main() {

	deferDemo1()

	deferDemo2Res := deferDemo2()
	fmt.Println("deferDemo2 res:", deferDemo2Res)

	deferDemo3Res := deferDemo3()
	fmt.Println("deferDemo3 res:", deferDemo3Res)

	deferDemo4Res := deferDemo4()
	fmt.Println("deferDemo4 res:", deferDemo4Res)

	deferDemo5Res := deferDemo5()
	fmt.Println("deferDemo5 res:", deferDemo5Res)
}

//多个defer的执行顺序是先加载后执行
func deferDemo1() {
	var i = 1
	defer fmt.Println("i1 =", i)
	i++
	defer func(i int) {
		fmt.Println("i2 =", i)
	}(i)
}

//defer 执行是在return 之后 函数结束之前
func deferDemo2() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

//提前定义好返回变量时 defer可以修改返回的变量
func deferDemo3() (result int) {
	defer func() {
		result++
	}()
	return 0
}

//return 的数据和函数返回的数据应该是一个值拷贝
func deferDemo4() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func deferDemo5() (r int) {
	defer func(r int) {
		r = r + 5 //此处r是一个局部变量
	}(r)
	return 1
}
