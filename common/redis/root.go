package redis

import (
	"sync"
)

var (
	redisApiInstance *RedisApi
	once             sync.Once
)

func GetRedisInstance(redisCfg *RedisCfg) *RedisApi {
	if redisCfg == nil {
		panic("Redis config is nil")
	}
	once.Do(func() {
		redisApiInstance = newRedis(redisCfg)
		if redisApiInstance == nil {
			panic("Can not init redis conn!!")
		}
	})

	return redisApiInstance
}

func UnInit() {
	if redisApiInstance != nil {
		redisApiInstance.redisClient.Close()
	}
}
