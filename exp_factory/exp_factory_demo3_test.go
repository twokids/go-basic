package exp_factory

import (
	"fmt"
	"testing"
)

func TestNewRedisCache(t *testing.T) {
	var redisCacheFactory CacheFactory1
	redisCacheFactory = RedisCacheFactory{}

	redisCache := redisCacheFactory.Create()

	redisCache.Set("nm", "12138")
	fmt.Println(redisCache.Get("nm"))
}
