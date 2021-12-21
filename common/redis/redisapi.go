package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	REDIS_TRUE  = 1
	REDIS_FALSE = 0
)

type RedisCfg struct {
	Address        string
	Password       string
	Index          int
	Prefix         string
	MaxIdle        int
	MaxActive      int
	IdleTimeout    int
	ReadTimeout    int
	WriteTimeout   int
	ConnectTimeout int
}

type RedisApi struct {
	redisClient *redis.Pool //redis连接池
	cfg         RedisCfg
}

func newRedis(cfg *RedisCfg) *RedisApi {

	redisServer := &RedisApi{
		redisClient: nil,
		cfg:         *cfg,
	}

	redisServer.redisClient = &redis.Pool{
		MaxIdle:     redisServer.cfg.MaxIdle + 1,
		MaxActive:   redisServer.cfg.MaxActive + 1,
		Wait:        true,
		IdleTimeout: time.Duration(redisServer.cfg.IdleTimeout) * time.Second,
		//PINGs connections that have been idle more than 30s
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Second*time.Duration(30) {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServer.cfg.Address,
				redis.DialDatabase(cfg.Index),
				redis.DialPassword(cfg.Password),
				redis.DialConnectTimeout(time.Duration(cfg.ConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(cfg.ReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(cfg.WriteTimeout)*time.Second))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	return redisServer
}

func (this *RedisApi) GetRedisClient() (redis.Conn, time.Time) {

	return this.getRedisClient()
}

func (this *RedisApi) GetRedisClientPool() (*redis.Pool) {
	return this.redisClient
}

//获取redis连接
func (this *RedisApi) getRedisClient() (redis.Conn, time.Time) {

	return this.redisClient.Get(), time.Now()
}

func (this *RedisApi) ReclaimRedisClient(conn redis.Conn, startTime time.Time) {
	this.reclaimRedisClient(conn, startTime)
}

func (this *RedisApi) reclaimRedisClient(conn redis.Conn, startTime time.Time) {
	if conn != nil {
		conn.Close()
	}
}

func (this *RedisApi) Exists(key string) (bool, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("EXISTS", key)
	if err != nil {
		return false, err
	}

	recData := v.(int64)
	return (recData == REDIS_TRUE), nil
}

func (this *RedisApi) Keys(predix string) ([]interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("KEYS", predix)
	if err != nil {
		return nil, err
	}

	if v != nil {
		return v.([]interface{}), nil
	}

	return nil, nil
}

func (this *RedisApi) Set(key string, value interface{}) error {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	_, err := rc.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (this *RedisApi) Incr(key string) error {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	_, err := rc.Do("INCR", key)
	if err != nil {
		return err
	}
	return nil
}

// 设置value，同时设置value过期时间(2.0.0版本后才支持)
func (this *RedisApi) SetEx(key string,
	expire int32, // 过期时间(秒)，-1不过期
	value interface{}) error {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	_, err := rc.Do("SETEX", key, expire, value)
	if err != nil {
		return err
	}
	return nil
}

func (this *RedisApi) SetNX(key string, value interface{}) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("SETNX", key, value)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *RedisApi) Expire(key string, value interface{}) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("EXPIRE", key, value)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *RedisApi) Get(key string) (value interface{}, err error) {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)
	v, err := rc.Do("GET", key)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *RedisApi) LLen(key string) (value int64, err error) {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)
	v, err := rc.Do("LLEN", key)
	if err != nil {
		return 0, err
	}
	return v.(int64), nil
}

func (this *RedisApi) RPush(key string, value ...interface{}) error {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	values := make([]interface{}, 0)
	values = append(values, key)
	values = append(values, value...)
	_, err := rc.Do("RPUSH", values...)
	if err != nil {
		return err
	}
	return nil
}

func (this *RedisApi) LPush(key string, value ...interface{}) error {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	values := make([]interface{}, 0)
	values = append(values, key)
	values = append(values, value...)
	_, err := rc.Do("LPUSH", values...)
	if err != nil {
		return err
	}
	return nil
}

func (this *RedisApi) RPop(key string) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("RPOP", key)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (this *RedisApi) BRPop(key string, timeout int64) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("BRPOP", key, timeout)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *RedisApi) GetRange(key string, start int64, end int64) ([]interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("LRANGE", key, start, end)
	if err != nil {
		return nil, err
	}

	recData := v.([]interface{})
	return recData, nil
}

func (this *RedisApi) LTrim(key string, start int64, end int64) error {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	_, err := rc.Do("LTRIM", key, start, end)
	if err != nil {
		return err
	}

	return nil
}

func (this *RedisApi) GetHashAllValues(key string) ([]interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("HVALS", key)
	if err != nil {
		return nil, err
	}

	recData := v.([]interface{})
	return recData, nil
}

func (this *RedisApi) GetHashAllKeys(key string) ([]interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("HKEYS", key)
	if err != nil {
		return nil, err
	}

	recData := v.([]interface{})
	return recData, nil
}

func (this *RedisApi) GetHashAll(key string) ([]interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("HGETALL", key)
	if err != nil {
		return nil, err
	}

	recData := v.([]interface{})
	return recData, nil
}

func (this *RedisApi) HSet(key string, field string, value interface{}) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("HSET", key, field, value)
	if err != nil {
		return nil, err
	}

	recData := v.(interface{})
	return recData, nil
}

func (this *RedisApi) PipeHSet(args ...interface{}) ([]interface{}, error) {
	if len(args)%3 != 0 {
		err := fmt.Errorf("PipeHSet args number is not times of 3")
		return nil, err
	}
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	rc.Send("MULTI")
	for i := 0; i < len(args); i = i + 3 {
		rc.Send("HSET", args[i], args[i+1], args[i+2])
	}
	recData, err := redis.Values(rc.Do("EXEC"))
	if err != nil {
		return nil, err
	}
	return recData, nil
}

func (this *RedisApi) HGet(key string, hkey string) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("HGET", key, hkey)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (this *RedisApi) PipeHDel(args ...interface{}) error {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	rc.Send("MULTI")
	for i := 0; i < len(args); i = i + 2 {
		rc.Send("HDEL", args[i], args[i+1])
	}
	_, err := redis.Values(rc.Do("EXEC"))
	if err != nil {
		return err
	}
	return nil
}

func (this *RedisApi) HExists(key string, field string) (bool, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("HEXISTS", key, field)
	if err != nil {
		return false, err
	}

	recData := v.(int64)
	return (recData == REDIS_TRUE), nil
}

func (this *RedisApi) GetSetAll(key string) ([]interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("SMEMBERS", key)
	if err != nil {
		return nil, err
	}

	recData := v.([]interface{})
	return recData, nil
}

func (this *RedisApi) IsSetMember(key string, member string) (bool, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("SISMEMBER", key, member)
	if err != nil {
		return false, err
	}

	recData := v.(int64)
	return (recData == REDIS_TRUE), nil
}

/*
func (this *RedisApi) PipeIsSetMember(args ...interface{}) ([]int, error) {

	if len(args)%2 != 0{
		err:= fmt.Errorf("PipeIsSetMember args number is not even")
		return nil, err
	}
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	rc.Send("MULTI")
	for i:=0; i< len(args); i=i+2  {
		rc.Send("SISMEMBER", args[i], args[i+1])
	}
	vs, err := rc.Do("EXEC")
	if err != nil {
		return nil, err
	}
	res:= make([]int, 0)
	for _,v :=range vs {
		iv, err := redis.Int(v, nil)
		if err != nil {
			return 0, keys, err
		}
		res = append(res, iv)
	}

	return res, nil
}
*/

func (this *RedisApi) HMGet(value ...interface{}) ([]interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("HMGET", value...)
	if err != nil {
		return nil, err
	}

	recData := v.([]interface{})
	return recData, nil
}

func (this *RedisApi) PipeHGet(args ...interface{}) ([]interface{}, error) {

	if len(args)%2 != 0 {
		err := fmt.Errorf("PipeHGet args number is not even")
		return nil, err
	}
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	rc.Send("MULTI")
	for i := 0; i < len(args); i = i + 2 {
		rc.Send("HGET", args[i], args[i+1])
	}
	recData, err := redis.Values(rc.Do("EXEC"))
	if err != nil {
		return nil, err
	}
	return recData, nil
}

func (this *RedisApi) HMSet(value ...interface{}) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("HMSET", value...)
	if err != nil {
		return nil, err
	}

	recData := v.(interface{})
	return recData, nil
}

func (this *RedisApi) SetSortSet(key string, args ...interface{}) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	values := make([]interface{}, 0)
	values = append(values, key)
	values = append(values, args...)
	v, err := rc.Do("ZADD", values...)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (this *RedisApi) GetSortSetRange(key string, start, end int64, withscore bool) ([]interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	var v interface{}
	var err error

	if withscore {
		v, err = rc.Do("ZRANGE", key, start, end, "WITHSCORES")
		if err != nil {
			return nil, err
		}
	} else {
		v, err = rc.Do("ZRANGE", key, start, end)
		if err != nil {
			return nil, err
		}
	}

	recData := v.([]interface{})
	return recData, nil
}

func (this *RedisApi) ZRem(key string, args ...interface{}) error {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	values := make([]interface{}, 0)
	values = append(values, key)
	values = append(values, args...)
	_, err := rc.Do("ZREM", values...)
	if err != nil {
		return err
	}

	return nil
}

//删除指定排序区间的元素
func (this *RedisApi) ZRemRangeByRank(key string, start, end int64) error {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	var err error

	_, err = rc.Do("ZREMRANGEBYRANK", key, start, end)
	if err != nil {
		return err
	}

	return nil
}

func (this *RedisApi) ZInterStore(dstKey string, args ...interface{}) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	values := make([]interface{}, 0)
	values = append(values, dstKey)
	values = append(values, args...)
	v, err := rc.Do("ZINTERSTORE", values...)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (this *RedisApi) ZCount(key string, min, max int) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("ZCOUNT", key, min, max)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (this *RedisApi) SetSet(key string, args ...interface{}) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	values := make([]interface{}, 0)
	values = append(values, key)
	values = append(values, args...)
	v, err := rc.Do("SADD", values...)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (this *RedisApi) GetSetMembers(key string) ([]interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("SMEMBERS", key)
	if err != nil {
		return nil, err
	}

	recData := v.([]interface{})
	return recData, nil
}

func (this *RedisApi) SInterStore(args ...interface{}) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("SINTERSTORE", args...)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (this *RedisApi) SDiffStore(args ...interface{}) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("SDIFFSTORE", args...)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (this *RedisApi) SScan(key string, args ...interface{}) (int, []string, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	keys := []string{}

	values := make([]interface{}, 0)
	values = append(values, key)
	values = append(values, args...)
	v, err := redis.Values(rc.Do("SSCAN", values...))
	if err != nil {
		return 0, keys, err
	}
	iter, err := redis.Int(v[0], nil)
	if err != nil {
		return 0, keys, err
	}
	k, err := redis.Strings(v[1], nil)
	if err != nil {
		return 0, keys, err
	}
	keys = append(keys, k...)

	return iter, keys, nil
}

func (this *RedisApi) SCard(key string) (interface{}, error) {

	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	v, err := rc.Do("SCARD", key)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (this *RedisApi) Pub(key string, period string, value interface{}) error {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)
	_, err := rc.Do("PUBLISH", key, value)
	if err != nil {
		return err
	}

	return nil
}

func (this *RedisApi) Delete(key interface{}) error {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)
	_, err := rc.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

func (this *RedisApi) TTL(key string) (interface{}, error) {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)
	v, err := rc.Do("TTL", key)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *RedisApi) PipeRPush(key string, value ...interface{}) (interface{}, error) {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	rc.Send("MULTI")
	for _, v := range value {
		rc.Send("RPUSH", key, v)
	}
	v, err := rc.Do("EXEC")
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *RedisApi) PipeRPushV2(args ...interface{}) (interface{}, error) {

	if len(args)%2 != 0 {
		err := fmt.Errorf("PipeRPushV2 args number is not times of 3")
		return nil, err
	}
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	rc.Send("MULTI")
	for i := 0; i < len(args); i = i + 2 {
		rc.Send("RPUSH", args[i], args[i+1])
	}

	v, err := rc.Do("EXEC")
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *RedisApi) PipeSetSet(key string, value ...interface{}) (interface{}, error) {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	rc.Send("MULTI")
	for _, v := range value {
		rc.Send("SADD", key, v)
	}
	v, err := rc.Do("EXEC")
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *RedisApi) PipeSetZSet(key string, value ...interface{}) (interface{}, error) {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	rc.Send("MULTI")
	for _, v := range value {
		rc.Send("ZADD", key, v)
	}
	v, err := rc.Do("EXEC")
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *RedisApi) Print() {

	fmt.Printf("%v/n", this.cfg)
}

type BatchResult struct {
	Reply interface{}
	Error error
}

//func (this *RedisApi) PipeBatchSet(itemMap map[string]interface{}) (result []BatchResult, err error) {
//	result = make([]BatchResult, 0)
//	rc, startTime := this.getRedisClient()
//	defer this.reclaimRedisClient(rc, startTime)
//
//	itemMapLen := len(itemMap)
//	for k, v := range itemMap {
//		sendErr := rc.Send("SET", k, v)
//		if sendErr != nil {
//			err = sendErr
//			return
//		}
//	}
//
//	flushErr := rc.Flush()
//	if flushErr != nil {
//		err = flushErr
//		return
//	}
//
//	for i := 0; i < itemMapLen; i++ {
//		ret, retErr := rc.Receive()
//		reply := ret
//		if retErr != nil {
//			reply = nil
//		}
//		result = append(result, BatchResult{Reply: reply, Error: retErr})
//	}
//
//	return
//}
//
//func (this *RedisApi) PipeBatchGet(keyList []string) (result []BatchResult, err error) {
//	result = make([]BatchResult, 0)
//	rc, startTime := this.getRedisClient()
//	defer this.reclaimRedisClient(rc, startTime)
//
//	listLen := len(keyList)
//	for i := 0; i < listLen; i++ {
//		sendErr := rc.Send("GET", keyList[i])
//		if sendErr != nil {
//			err = sendErr
//			return
//		}
//	}
//
//	flushErr := rc.Flush()
//	if flushErr != nil {
//		err = flushErr
//		return
//	}
//
//	for i := 0; i < listLen; i++ {
//		ret, retErr := rc.Receive()
//		reply := ret
//		//if retErr != nil {
//		//	reply = nil
//		//}
//		result = append(result, BatchResult{Reply: reply, Error: retErr})
//	}
//
//	return
//}

func (this *RedisApi) PipeBatchExpire(keyList []string, value interface{}) (result []BatchResult, err error) {
	result = make([]BatchResult, 0)
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)

	listLen := len(keyList)
	for i := 0; i < listLen; i++ {
		rc.Send("EXPIRE", keyList[i], value)
	}

	flushErr := rc.Flush()
	if flushErr != nil {
		err = flushErr
		return
	}

	for i := 0; i < listLen; i++ {
		ret, retErr := rc.Receive()
		reply := ret
		if retErr != nil {
			reply = nil
		}
		result = append(result, BatchResult{Reply: reply, Error: retErr})
	}

	return
}

/*	注意值不能太多，测试50w是ok，100w不ok
	经测试效率比pipeline效率高不少
*/
func (this *RedisApi) MSet(itemMap map[string]interface{}) (reply interface{}, err error) {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)
	params := make([]interface{}, 0)
	for k, v := range itemMap {
		params = append(params, k)
		params = append(params, v)
	}

	v, err := rc.Do("MSET", params...)
	if err != nil {
		return nil, err
	}

	recData := v.(interface{})
	return recData, nil
}

func (this *RedisApi) MGet(keyList []string) (reply interface{}, err error) {
	rc, startTime := this.getRedisClient()
	defer this.reclaimRedisClient(rc, startTime)
	params := make([]interface{}, 0)
	listLen := len(keyList)
	for i := 0; i < listLen; i++ {
		params = append(params, keyList[i])
	}

	v, err := rc.Do("MGET", params...)
	if err != nil {
		return nil, err
	}

	recData := v.(interface{})
	return recData, nil
}
