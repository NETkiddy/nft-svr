package redis

import (
	"fmt"
	"github.com/NETkiddy/common-go/svr_adapter/glue/command"
	"testing"

	"github.com/NETkiddy/common-go/config"
	"github.com/gomodule/redigo/redis"
)

var server = &Server{}

func Test_GetRedisInstance(t *testing.T) {
	//注册配置文件
	command.Init("test-redis", server)
	config.GetConfigManager().Register("app")
	command.Execute()

	appConfig := config.GetViper("app")
	redisCfg := &RedisCfg{
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
	c := GetRedisInstance(redisCfg)
	if c == nil {
		fmt.Println("GetRedisInstance failed")
		return
	}

	v, err := c.RPop("test_bpop")
	if err != nil {
		fmt.Println(err)
		return
	}
	if v == nil {
		fmt.Println("test_bpop is empty")
	}
	data, err := redis.ByteSlices(v, err)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("data:", data)

	/*
		r, err := c.ZInterStore("z2", 1, "s1", "weights", 2)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(r)
	*/
	/*
		values := []interface{}{1, "a", 3, "b", 2, "c"}
		r, err := c.SetSortSet("test_zadd", values...)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(r)
	*/
	/*
		values := []interface{}{"a", "b", "c"}
		r, err := c.SetSet("test_sadd", values...)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(r)
	*/
	/*
		iter := 0
		i := 0
		var r []string
		var err error
		for {
			iter, r, err = c.SScan("test_sscan", iter)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(iter, "==", r)
			if iter == 0 {
				fmt.Printf("done read redis [test_sscan]\n")
				break
			}
			if i > 10 {
				break
			}
			i++
		}
	*/
	/*
			err := c.Set("audience:statistic", `"{
		   "data":[
		      {
		         "_index":"cge",
		         "_type":"audience",
		         "_id":"8",
		         "_score":1.2039728,
		         "_source":{
		            "age":"54",
		            "city":"济南",
		            "gender":"女",
		            "mobile":"13800138007",
		            "name":"张小泉",
		            "privince":"山东"
		         }
		      },
		      {
		         "_index":"cge",
		         "_type":"audience",
		         "_id":"9",
		         "_score":1.2039728,
		         "_source":{
		            "age":"18",
		            "city":"济南",
		            "gender":"女",
		            "mobile":"13800138008",
		            "name":"赵括",
		            "privince":"山东"
		         }
		      }
		   ]
		}"`)
			if err != nil {
				fmt.Println(err)
				return
			}

			is := []string{"a", "va", "b", "vb"}
			s := make([]interface{}, 0, len(is))
			s = append(s, "SSS")
			for _, v := range is {
				s = append(s, v)
			}
			r, err := c.HMSet(s...)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(r)

			ixxx := []int{1, 2, 3, 4}
			xxx := make([]interface{}, 0, len(ixxx))
			for _, v := range ixxx {
				xxx = append(xxx, v)
			}
			r, err = c.PipeSetSet("YYY", xxx...)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(r)
	*/
	/*
		r, err = c.PipeRPush("XXX", xxx...)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(r)
		v, err := c.SInterStore("Y", "A", "B", "C")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)
		vs, err := c.GetSetMembers("Y")
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, vv := range vs {
			vvv, ok := vv.([]uint8)
			if !ok {
				fmt.Println("not ok")
			}
			fmt.Println(string(vvv))
		}
		v, err = c.SCard("Y")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)
	*/
}

type Server struct {
}

func (this *Server) Start() {
}
func (this *Server) Stop() {
}
