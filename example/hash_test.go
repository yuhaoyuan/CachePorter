package example

import (
	"CachePorter/RedisHandle"
	"fmt"
	"testing"
)

func computingFunc(param ...interface{}) (interface{}, error) {
	rs := make(map[string]string)
	for _, param := range param {
		rs[param.(string)] = "1"
	}
	return rs, nil
}

func TestCustomString(t *testing.T) {
	uid := "10001"
	key := "test-key"
	redisCli := NewRedisClient()
	defer func() {
		redisCli.Del(key)
	}()

	rPorter := RedisHandle.NewRedisPorter(key, RedisHandle.HashKey, redisCli, RedisHandle.DefaultOptions(), RedisHandle.ComputingFuncOption([]interface{}{uid}, computingFunc), RedisHandle.ReadingFuncOption([]interface{}{RedisHandle.HGet, uid}))

	err := rPorter.Read(uid) // computingFunc's param
	if err != nil {
		fmt.Println("err = ", err)
	}

	rs, err := redisCli.Get("key").Result()
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("rs = ", rs)
}
