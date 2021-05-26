package exp_routine

import (
	"fmt"
	"sync"
	"time"
)

var memberNameChan chan string = make(chan string)
var memberCountChan chan int = make(chan int)
var wg sync.WaitGroup

func Goroutine_main() {
	//fmt.Println(runtime.NumCPU())
	//runtime.GOMAXPROCS(runtime.NumCPU() - 1) //设置cpu的最大核心数

	//go关键字，并发操作
	//go loop()
	//go loop()

	//go send()
	//go receive()

	//readfromtxt()
	//go writetosql()
	//wg.Wait()
	//fmt.Println("all done")

	//chan_demo3()
	//time.Sleep(time.Second * 3)
}

func loop() {
	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond * 5)
		fmt.Printf("%d,", i)
	}
}

func send() {
	memberNameChan <- "lisa"
	memberCountChan <- 1
	time.Sleep(time.Microsecond * 5)
	memberNameChan <- "lucy"
	memberCountChan <- 2
	time.Sleep(time.Microsecond * 5)
	memberNameChan <- "lily"
	memberCountChan <- 3
	time.Sleep(time.Microsecond * 5)
}

func receive() {
	//memberName := <-memberNameChan
	//fmt.Println(memberName)
	//memberCount := <-memberCountChan
	//fmt.Println(memberCount)
	//memberName = <-memberNameChan
	//fmt.Println(memberName)
	//memberCount = <-memberCountChan
	//fmt.Println(memberCount)
	//memberName = <-memberNameChan
	//fmt.Println(memberName)
	//memberCount = <-memberCountChan
	//fmt.Println(memberCount)

	for {
		select {
		case memberName := <-memberNameChan:
			fmt.Println(memberName)
		case memberCount := <-memberCountChan:
			fmt.Println(memberCount)

		}
	}
}

//从txt中读取内容，写入协程
func readfromtxt() {
	for i := 0; i < 3; i++ {
		//从txt中读取数据
		wg.Add(1)
	}
}

//从协程读取，写入sql
func writetosql() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Microsecond * 5)
		fmt.Println("done->", i)
		wg.Done()

		//再写入sql中
	}
}


func add(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1
}

func chan_demo3() {
	start := time.Now()
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add(1, i, chs[i])
	}
	for _, ch := range chs {
		fmt.Println("for")
		<- ch
	}
	end := time.Now()
	consume := end.Sub(start).Microseconds()
	fmt.Println("程序执行耗时(ms)：", consume)
}