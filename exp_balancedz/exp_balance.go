package exp_balancedz

import (
	"fmt"
	bl "go-basic/exp_balancedz/util/balance"
	"math/rand"
	"time"
)

func Balance_main() {

	ab_simple_balance()

	ab_balance()

	dz_balance()
}

func ab_balance()  {
	//控制AB流量百分比

	//设定2种分配方式的比例为70：30
	bili:=70
	//验证其对应比例
	ab_result:=make(map[int]int)
	ab_result[0]=0//随机分配
	ab_result[1]=0//等级随机

	abstart:=time.Now()
	var ablst=make([]int,100)

	//ablst:=[100]int{}
	for i:=0;i<100;i++{
		ablst[i]=i
	}
	for j:=0;j<10000;j++{
		inst, err := bl.DoBalance2("roundrobin", ablst)
		if err != nil {
			fmt.Println("no balance err", err)
			continue
		}
		if inst<bili{
			ab_result[0]++
		}else{
			ab_result[1]++
		}
		//time.Sleep(time.Second)
	}
	abdur:=time.Since(abstart)
	fmt.Println("------abdur-----",abdur)
	fmt.Println("dz list ,",ab_result)
	fmt.Println("abdur end")
}

func ab_simple_balance()  {
	//控制5，5百分比流量
	//1->随机分配,2->等级随机分配
	ablst:=[]int{1,2}

	abrst:=make(map[int]int)//统计ab的分配次数
	abstart:=time.Now()
	for j:=0;j<10000;j++{
		inst, err := bl.DoBalance2("roundrobin", ablst)
		if err != nil {
			fmt.Println("no balance err", err)
			continue
		}
		fmt.Println(inst)

		abrst[inst]++
		//time.Sleep(time.Second)
	}
	abdur:=time.Since(abstart)
	fmt.Println("------abdur-----",abdur)
	fmt.Println("dz list ,",abrst)
	fmt.Println("abdur end")
}

func dz_balance()  {
	//控制随机分配店长
	var dzinsts []*bl.Instance

	rst:=make(map[string]int)
	for i := 0; i < 8; i++ {
		name := fmt.Sprintf("兔子%d号",i+1)
		uaid := rand.Intn(255)
		one := bl.NewInstance(name, uaid)
		dzinsts = append(dzinsts, one)
		rst[name]=0
	}

	var balanceName = "hash"//hash random roundrobin
	//if len(os.Args) > 1 {
	//	balanceName = os.Args[1]
	//}
	start:=time.Now()
	for j:=0;j<10000;j++{
		inst, err := bl.DoBalance(balanceName, dzinsts)
		if err != nil {
			fmt.Println("no balance err", err)
			continue
		}
		fmt.Println(inst)

		rst[inst.GetName()]++
		//time.Sleep(time.Second)
	}
	dur:=time.Since(start)
	fmt.Println("------dur-----",dur)
	fmt.Println("dz list ,",rst)
	fmt.Println("end")
}
