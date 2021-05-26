package main

import (
	"fmt"
	exp03 "go-basic/exp_json"
	exp02 "go-basic/exp_point"
	exp01 "go-basic/exp_routine"
)

func main() {
	fmt.Println("this is main")
	arrayslice_main()
	make_main()
	routines_main()
	new_main()
	appendcopydelete_main()
	panic_recover_main()
	lencapclose_main()
	//struct_main()
	exp01.Goroutine_main()
	exp02.Point_main()
	exp03.Json_main()

}
