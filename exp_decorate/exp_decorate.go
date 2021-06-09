package exp_decorate

import "fmt"

//装饰模式
type Component interface {
	Operate()
}

type ComponentObj struct {
}

func (c *ComponentObj) Operate() {
	fmt.Println("实现抽象接口")
}

//定义抽象的装饰者
type Decorator interface {
	Component
	Do() //这个是额外的方法
}

//实现一个具体的装饰者
type DecoratorObj struct {
	c Component
}

func (d *DecoratorObj) Operate() {
	d.Do()        //额外补充的行为
	d.c.Operate() //原先的装饰行为
}

func (d *DecoratorObj) Do() {
	fmt.Println("装饰行为")
}
