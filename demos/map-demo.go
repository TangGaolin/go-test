package main

import "fmt"

func main() {
	mapStcTest()
}

func mapStcTest() {

	type Student struct {
		Name string
		Id   int
	}

	mapStu := map[string]Student{}
	mapStu["demo"] = Student{
		Name: "demo",
		Id:   111,
	}
	//mapStu["demo"].Id = 222   // 这种写法会build error
	//在golang中, map中的值是不可以原地修改的，go中map中的赋值属于值copy。

	//那么，如何在go语言中原地修改map中的value呢？ 答案是：传指针！
	mapStu2 := map[string]*Student{}
	mapStu2["demo"] = &Student{
		Name: "demo",
		Id:   111,
	}
	mapStu2["demo"].Id = 222
	fmt.Println(mapStu2["demo"])
}
