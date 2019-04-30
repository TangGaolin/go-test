package main

import (
	"fmt"
	"os"
	"time"
)
/******
panic会停掉当前正在执行的程序，不只是当前协程。
在这之前，它会有序地执行完当前协程defer列表里的语句，
其它协程里挂的defer语句不作保证。
因此，我们经常在defer里挂一个recover语句，防止程序直接挂掉，
这起到了 try...catch的效果。

注意，recover()函数只在defer的上下文中才有效（且只有通过在defer中用匿名函数调用才有效），直接调用的话，只会返回 nil.
******/

func main()  {
	defer fmt.Println("defer main")
	var user= os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err: ", err)
			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here ... ")
			}()
			if user == "" {
				panic("should set user env.")
			}
			fmt.Println("after panic...")
		}()
	}()

	time.Sleep(time.Second)
	fmt.Println("end of main function.")
}


