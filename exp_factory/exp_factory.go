package exp_factory

import "fmt"

type Animal interface {
	SetAge(age int32) int32
	GetAge() int32
}

type Cat struct {
	Age int32
}

func (c *Cat) SetAge(age int32) int32 {
	c.Age = age
	fmt.Printf("cat age", c.Age)
	return c.Age
}
func (c *Cat) GetAge() int32 {

	fmt.Printf("cat age", c.Age)
	return c.Age
}

type Dog struct {
	Age int32
}

func (d *Dog) SetAge(age int32) int32 {
	d.Age = age
	fmt.Printf("dog age", d.Age)
	return d.Age
}
func (d *Dog) GetAge() int32 {
	fmt.Printf("dog age", d.Age)
	return d.Age
}

type AnimalType int

const (
	catType AnimalType = iota //0
	dogType                   //1
)

type AnimalAgeFactory struct {
}

func (a AnimalAgeFactory) Create(animalType AnimalType) Animal {
	if animalType == catType {
		return &Cat{}
	} else if animalType == dogType {
		return &Dog{}
	}
	return nil

}
