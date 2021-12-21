package common

import (
	"errors"
	"github.com/NETkiddy/common-go/config"
	"github.com/NETkiddy/nft-svr/common/redis"
)

func GetRedisClient() (c *redis.RedisApi, err error) {
	appConfig := config.GetViper("app")
	redisCfg := &redis.RedisCfg{
		Address:        appConfig.GetString("redis.addr"),
		Password:       appConfig.GetString("redis.password"),
		Index:          appConfig.GetInt("redis.index"),
		Prefix:         appConfig.GetString("redis.prefix"),
		MaxIdle:        appConfig.GetInt("redis.maxidle"),
		MaxActive:      appConfig.GetInt("redis.maxactive"),
		IdleTimeout:    appConfig.GetInt("redis.idletimeout"),
		ReadTimeout:    appConfig.GetInt("redis.readtimeout"),
		WriteTimeout:   appConfig.GetInt("redis.writetimeout"),
		ConnectTimeout: appConfig.GetInt("redis.connecttimeout"),
	}

	c = redis.GetRedisInstance(redisCfg)
	if c == nil {
		return c, errors.New("Get redis connection failed.")
	}

	return c, nil
}
