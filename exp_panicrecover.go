package main

import (
	"fmt"
)

func panic_recover_main() {
	//receivePanic1()
	//receivePanic2()
}

func receivePanic1() {
	//接收异常panic
	defer func() {
		message := recover()
		fmt.Println("panic message:", message)
	}()
	//抛出错误
	panic("I am panic")

}

func receivePanic2() {
	//接收异常panic
	defer coverPanic2()
	//抛出错误
	//panic(errors.New("I am a error panic"))
	panic(404)

}

func coverPanic2() {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string message:",message)
	case error:
		fmt.Println("error message:",message)
	default:
		fmt.Println("unknow message:",message)

		
	
	}
}
