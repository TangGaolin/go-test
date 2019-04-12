package main

import (
	"fmt"
	"unsafe"
)

type A int64

func convertIntArrToA(x []int64) []A {
	return *(*[]A)(unsafe.Pointer(&x))
}

func byteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// Go语言中类型转换
func main() {
	x := []int64{1, 2, 3}
	fmt.Println(convertIntArrToA(x))

	str := "hello world!"
	fmt.Println(byteSlice2String([]byte(str)))
}
