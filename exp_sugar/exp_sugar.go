package exp_sugar

import "fmt"

/*
常用语法糖
1，...
2, :=
*/
func Sugar_main() {
	//sugar_demo("22", "11", "33")
}

func sugar_demo(values ...string) {
	for key, value := range values {
		fmt.Println(key, value)
	}

	name := "lisa"
	fmt.Println(name)
}
