package exp_point

import "fmt"

func Point_main()  {
	//point_demo()
}

func point_demo()  {
	var count int=30
	var countpoint *int
	countpoint=&count

	fmt.Printf("count address %x \n",&count) //变量的地址
	fmt.Printf("countpoint address %x \n",countpoint) //变量存储的地址
	fmt.Printf("count value %d \n",count)
	fmt.Printf("countpoint value %d \n",*countpoint) //指针，指向地址的值
}
