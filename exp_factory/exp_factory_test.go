package exp_factory

import (
	"fmt"
	"testing"
)

func Test_Anmial_01(t *testing.T) {
	var kitty = new(Cat)
	kitty.SetAge(2)
	fmt.Println("kitty age", kitty.GetAge())

	var dahuang = &Dog{}
	dahuang.SetAge(3)
	fmt.Println("dahuang age", dahuang.GetAge())
}

func Test_Anmial_02(t *testing.T) {
	var af = new(AnimalAgeFactory)

	cur_type := AnimalType(dogType)
	switch cur_type {
	case dogType:
		xiaohei := af.Create(dogType)
		xiaohei.SetAge(200)
		fmt.Println("xiaohei age", xiaohei.GetAge())
	case catType:
		dali := af.Create(catType)
		dali.SetAge(5)
		fmt.Println("dali age", dali.GetAge())
	}
}
