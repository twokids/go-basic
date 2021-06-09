package exp_factory

import "errors"

type Cache interface {
	Set(key string, value string)
	Get(key string) (value string)
}

type CacheType int

const (
	RedisType CacheType = iota
	MemType
)

//redis
type RedisCache struct {
	data map[string]string
}

func (i *RedisCache) Set(key string, value string) {
	i.data[key] = value
}

func (i *RedisCache) Get(key string) (value string) {
	return i.data[key]
}

//mem
type MemCache struct {
	data map[string]string
}

func (m MemCache) Set(key string, value string) {
	m.data[key] = value
}

func (m MemCache) Get(key string) (value string) {
	return m.data[key]
}

type CacheFactory struct {
}

func (c *CacheFactory) Create(cacheType CacheType) (Cache, error) {
	switch cacheType {
	case RedisType:
		return &RedisCache{
			data: map[string]string{},
		}, nil
	case MemType:
		return &MemCache{
			data: map[string]string{},
		}, nil

	}
	return nil, errors.New("error cache type")
}
