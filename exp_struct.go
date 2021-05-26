package main

import "fmt"

type Behavior interface {
	Eat() string
	Run() string
}

type Animal struct {
	Colour string
}

func (a *Animal) Eat() string {
	fmt.Println(a.Colour,"eat eat eat~~")
	return "eat"
}

//Dog隐式的继承了Behavior
type Dog struct {
	Animal //继承父类Animal
	ID     int
	Name   string
	Age    int
}

func (d *Dog) Run() string {
	fmt.Println(d.Name, "dog is running~")
	return "run"
}

type Cat struct {
	Animal //继承父类Animal
	Name   string
}

func (c *Cat) Run() string {
	fmt.Println(c.Name, "cat is running~")
	return "run"
}

func action(b Behavior) string {
	b.Run()
	b.Eat()
	return ""
}

func struct_main() {
	var dog Dog
	dog.ID = 1
	dog.Name = "lily"
	dog.Age = 2
	fmt.Println("dog:", dog)
	dog.Run() //如果是外部调用，run方法首字母需要大写，小写只能在包内引用
	dog.Eat()

	cat := Cat{
		Name: "lucy",
	}
	fmt.Println("cat:", cat)
	cat.Run()
	cat.Eat()

	dog2 := new(Dog)
	dog2.ID = 3
	dog2.Name = "lilei"
	dog2.Age = 4
	fmt.Println("dog2:", dog2) //dog2是指针，指针可以当引用可以点出其属性值

	var animal001 Behavior
	animal001 = new(Cat)
	animal001.Run()
	animal001.Eat()

	dog01:=new(Dog)
	cat01:=new(Cat)
	action(dog01)
	action(cat01)
}
