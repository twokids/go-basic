package exp_balancedz

import (
	"errors"
)

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})
}

type RoundRobinBalance struct {
	curIndex int
	curAbIndex int
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instence")
		return
	}

	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}

	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}

func  (p *RoundRobinBalance)DoBalance2(insts []int,  key ...string) (inst int, err error)  {
	if len(insts) == 0 {
		err = errors.New("No instence")
		return
	}

	lens := len(insts)
	if p.curAbIndex >= lens {
		p.curAbIndex = 0
	}

	inst = insts[p.curAbIndex]
	p.curAbIndex = (p.curAbIndex + 1) % lens
	return
}