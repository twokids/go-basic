package exp_factory

//定义抽象的cache的cache工厂
type CacheFactory1 interface {
	Create() Cache
}

func NewRedisCache() *RedisCache {
	return &RedisCache{data: map[string]string{}}
}

func NewMemCache() *MemCache {
	return &MemCache{data: map[string]string{}}
}

//实现具体的工厂：redis工厂
type RedisCacheFactory struct {
}

func (rf RedisCacheFactory) Create() Cache {
	return NewRedisCache()
}

//实现具体的工厂：mem工厂
type MemCacheFactory struct {
}

func (mf MemCacheFactory) Create() Cache {
	return NewMemCache()
}
