package exp_balancedz

type Balancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)


	DoBalance2([]int, ...string) (int, error)
}
