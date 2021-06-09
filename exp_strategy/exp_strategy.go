package exp_strategy

import "fmt"

//策略模式
type Context struct {
	Strategy
}

//抽象的策略
type Strategy interface {
	Do()
}

//实现具体的策略
type Strategy1 struct {
}

func (s1 Strategy1) Do() {
	fmt.Println("s1 do")
}

//实现具体的策略
type Strategy2 struct {
}

func (s2 Strategy2) Do() {
	fmt.Println("s2 do")
}
