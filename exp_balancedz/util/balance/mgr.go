package exp_balancedz

import "fmt"

type BalanceMgr struct {
	allBalancer map[string]Balancer
}

var mgr = BalanceMgr{
	allBalancer: make(map[string]Balancer),
}

func (p *BalanceMgr) registerBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balance", name)
		return
	}

	fmt.Printf("use %s balancer\n", name)
	inst, err = balancer.DoBalance(insts)
	return
}


func DoBalance2(name string, insts []int) (inst int, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balance", name)
		return
	}

	fmt.Printf("use %s balancer\n", name)
	inst, err = balancer.DoBalance2(insts)
	return
}
/*DoBalance([]*Instance, ...string) (*Instance, error)


DoBalance2([]int, ...string) (int, error)*/