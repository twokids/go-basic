package exp_decorate

import "testing"

func TestComponentObj_Operate(t *testing.T) {
	c := &ComponentObj{}
	c.Operate()
}

func TestDecoratorObj_Operate(t *testing.T) {
	d := &DecoratorObj{}
	c := &ComponentObj{}
	d.c = c
	//d.Do()
	d.Operate()

}
