package exp_balancedz

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instence")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]

	return
}

func  (p *RandomBalance)DoBalance2(insts []int,  key ...string) (inst int, err error)  {
	return 1,nil
}