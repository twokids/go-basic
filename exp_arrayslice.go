package main

import "fmt"

//数组
func initArray() {
	var arrayInt64 [3]int64
	arrayInt64[0], arrayInt64[1], arrayInt64[2] = 0, 4, 3
	fmt.Println(arrayInt64)

	arrayString:=[]string{"hello","go"}
	//arrayString[4]="" 超出边界赋值，编写无误，运行会报错
	fmt.Println(arrayString)

	arrayFloat:=[...]float64{1.0,2.3,4.6}
	//arrayFloat[4]=9 超出边界赋值，有...，编码即检查
	fmt.Println(arrayFloat)

	//  ...可以加在第一个[]中，第二个不允许
	matrix:=[...][]int64{
		{1,2},
		{2,3},
	}
	fmt.Println(matrix)
	fmt.Println(matrix[0][1])
	fmt.Println(matrix[0])
}

func initSlice()  {
	/*
		s[n]    切片s中索引为n的项
		s[n:m]  从切片s的索引位置n到m-1处获得的切片
		s[n:]	从切片s的所有位置n到len(s)-1处
		s[:m]	从切片s的所有位置0到m-1处
		s[:]	从切片s的所有位置len(s)-1处
		cap(s)	切片s的容量：总是大于等于len(s)
		len(s)	切片s包含的个数
		s[:cap(s)] 增加切片的长度从0到其容量，两者不同有效
	*/



}

func arrayslice_main() {
	//fmt.Println(1)
	//initArray()
}
