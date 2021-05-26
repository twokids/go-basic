package main

import "fmt"

//len string array slice map chan
//cap slice array chan
//close chan

func lencapclose_main()  {
	//getLen()
	//closeChan()
}


func getLen()  {
	mIDSliceDst:=make([]string,3,5)
	mIDSliceDst[0]="11"
	mIDSliceDst[1]="22"
	mIDSliceDst[2]="33"
	fmt.Println(len(mIDSliceDst),cap(mIDSliceDst))
	mIDSliceDst=append(mIDSliceDst,"44")
	mIDSliceDst=append(mIDSliceDst,"44")
	mIDSliceDst=append(mIDSliceDst,"44")
	fmt.Println(len(mIDSliceDst),cap(mIDSliceDst))
}

func closeChan()  {
	mChan:=make(chan int,1)
	close(mChan)//defer表示整个函数结束后执行
	mChan<-1 //往chan中写入数据  panic: send on closed channel
	//业务代码
}