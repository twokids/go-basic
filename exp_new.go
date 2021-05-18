package main

import (
	"fmt"
	"reflect"
)

//返回传入类型的指针地址
//内存置零
func NewMap() {
	mNewMap := new(map[int]string)
	fmt.Println("new:", reflect.TypeOf(mNewMap)) //*map[int]string
	mMakeMap := make(map[int]string)
	fmt.Println("make:", reflect.TypeOf(mMakeMap)) //map[int]string

}

func new_main() {
	//NewMap()
}
