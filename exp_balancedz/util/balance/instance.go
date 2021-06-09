package exp_balancedz

import "strconv"

type Instance struct {
	name string
	uaid int
}

func NewInstance(name string, uaid int) *Instance {
	return &Instance{
		name: name,
		uaid: uaid,
	}
}

func (p *Instance) GetName() string {
	return p.name
}

func (p *Instance) GetUaid() int {
	return p.uaid
}

func (p *Instance) String() string {
	return p.name + ":" + strconv.Itoa(p.uaid)
}
