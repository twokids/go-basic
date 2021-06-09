package exp_balancedz

import (
	"fmt"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
	key string
}

func init() {
	RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]

	return
}

func  (p *HashBalance)DoBalance2(insts []int,  key ...string) (inst int, err error)  {
	return 1,nil
}