package main

import "fmt"

// 返回的3种引用类型
func make_main()  {
	//makeSlice()
	//makeMap()
	//makeChan()
}

func makeSlice()  {
	mSlice:=make([]string,3)
	mSlice[0]="dog"
	mSlice[1]="cat"
	mSlice[2]="tiger"
	fmt.Println(mSlice)
}

func makeMap()  {
	mMap:=make(map[int]string)
	mMap[8]="lisa"
	mMap[32]="chb"
	fmt.Println(mMap)
}

func makeChan()  {
	mChan:=make(chan int)
	close(mChan)
}