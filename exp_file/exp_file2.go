package exp_file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type midList struct {
	mid []string
}

func File2_Main() {
	fmt.Println("script begin.... ", time.Now().Format("2006-01-02 15:04:05"))
	timeObj := time.Now()
	yesTime := timeObj.AddDate(0, 0, -1)
	yesterday := yesTime.Format("20060102")
	mid_file := "/data/dislike/triggerdislike/" + yesterday + "/uniq"

	mid_file="D:\\test\\D373_unionId_202107051018341.txt"


	read_start:=time.Now()
	fp, err := os.Open(mid_file)
	defer fp.Close()
	if err != nil {
		fmt.Println(mid_file, err)
		return
	}

	var contents []midList
	var content midList
	buf := bufio.NewReader(fp)
	iterator := 0
	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				//                fmt.Println("meet file ending")
			} else {
				fmt.Println(err)
			}
			break
		}
		iterator++

		mid := strings.TrimSpace(line)
		//将文件切割成每2000行一份
		if iterator <= 2000 {
			content.mid = append(content.mid, mid)
		} else {
			iterator = 1
			contents = append(contents, content)
			content.mid = nil
			content.mid = append(content.mid, mid)
		}

	}
	contents = append(contents, content)
	//
	read_dur:=time.Since(read_start)
	fmt.Println("------read_dur-----",read_dur)

	for _, mid_list := range contents {
		wg.Add(1) //这里也可以直接使用chan来完成通信
		go mid_list.MidToDo()
	}
	wg.Wait()

	fmt.Println("script end.... ", time.Now().Format("2006-01-02 15:04:05"))
}

func (i midList) MidToDo() {
	//TODO
}
