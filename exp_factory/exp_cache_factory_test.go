package exp_factory

import (
	"fmt"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cf:=&CacheFactory{}


	redisC,error:=cf.Create(RedisType)
	if error != nil {
		t.Error(error)
	}
	redisC.Set("name","lisa")
	fmt.Println("redis cache,",redisC.Get("name"))


	memC,error:=cf.Create(MemType)
	if error != nil {
		t.Error(error)
	}
	memC.Set("name","lily")
	fmt.Println("mem cache,",memC.Get("name"))

}
