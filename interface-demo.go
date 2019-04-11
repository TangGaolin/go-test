package main

import (
	"fmt"
)

type Pet interface {
	SayName()
}

type Dog struct {
	Name string
}

func (d *Dog) SayName() {
	fmt.Println(d.Name)
}

func main() {
	var dog1 *Dog
	if dog1 == nil {
		fmt.Println("The first dog is nil.")
	}

	dog2 := dog1
	if dog2 == nil {
		fmt.Println("The second dog is nil.")
	}

	var pet Pet = dog2
	// pet != nil
	// 当我们在给一个接口变量赋值的时候，该变量的动态类型会与它的动态值一起被存储在一个专用的数据结构中。
	// 所以此时的pet 包含了一个指针 指向类型信息 另外一直指针指向动态值也就是nil
	// 这种数据接口在go语言中的runtime包中就叫 iface
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}

	dog2 = &Dog{}
	pet = dog2
	dog2.Name = "二哈"
	pet.SayName()

}
