package main

import "fmt"

func appendcopydelete_main()  {
	//appendEleForSlice()
	//copyEleForSlice()
	//deleteFromMap()
}

func appendEleForSlice()  {
	mIDSlice:=make([]string,2)
	mIDSlice[0]="11"
	mIDSlice[1]="22"
	//mIDSlice[2]="233" 出错
	mIDSlice=append(mIDSlice,"33") //长度+1 cap容量+2，长度超出容量时，会翻倍扩容(不完全是)，如果容量太大不会再翻倍
	//fmt.Println(mIDSlice)  该行注释后，上下两行执行后，内存和长度一起增加到4，没有len3 cap4的过程了
	mIDSlice=append(mIDSlice,"44") //len 4 cap 4
	fmt.Println(mIDSlice)
}

func copyEleForSlice()  {
	mIDSlice1:=make([]string,2)
	mIDSlice1[0]="111"
	mIDSlice1[1]="122"

	mIDSlice2:=make([]string,3)
	mIDSlice2[0]="211"
	mIDSlice2[1]="222"
	mIDSlice2[2]="233"

	copy(mIDSlice1,mIDSlice2)				//copy不影响2个切片的长度，不会动态扩容，仅在切片1的有限长度上，把切片2 对应的位置复制到切片1上
	fmt.Println("copy 1:",mIDSlice1)	//[211 222]
	fmt.Println("copy 2:",mIDSlice2)	//[211 222 233]
}

func deleteFromMap()  {
	mIDMap:=make(map[int]string)
	mIDMap[0]="11"
	mIDMap[1]="22"
	fmt.Println(mIDMap,len(mIDMap))
	delete(mIDMap,0)
	delete(mIDMap,1)
	fmt.Println(mIDMap)
}