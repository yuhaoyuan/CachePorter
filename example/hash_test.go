package example

import (
	"CachePorter/RedisHandle"
	"CachePorter/RedisHandle/service"
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

func TestHGet(t *testing.T) {
	uid := "10001"
	key := "test-key"
	redisCli := InitMiniRedisClient()
	defer func() {
		redisCli.Del(key)
	}()

	readingParam := []interface{}{redisHandle.HGet, uid}
	computingParam := []interface{}{uid}
	rPorter := service.NewRedisPorter(key, redisHandle.HashKey, redisCli, service.DefaultOptions(), service.ComputingFuncOption(computingParam, computingFunc), service.ReadingFuncOption(readingParam))

	data, err := rPorter.Read()
	if err != nil {
		fmt.Println("err = ", err)
	}

	rs, err := redisCli.HGet("key", uid).Result()
	if err != nil {
		fmt.Println("err = ", err)
	}

	if rs != data {

	}

}

func TestHMGet(t *testing.T) {
	uid := "10001"
	key := "test-key"
	redisCli := InitMiniRedisClient()
	defer func() {
		redisCli.Del(key)
	}()

	readingParam := []interface{}{redisHandle.HMGet, uid}
	computingParam := []interface{}{uid}
	rPorter := service.NewRedisPorter(key, redisHandle.HashKey, redisCli, service.DefaultOptions(), service.ComputingFuncOption(computingParam, computingFunc), service.ReadingFuncOption(readingParam))

	data, err := rPorter.Read()
	if err != nil {
		fmt.Println("err = ", err)
	}

	rs, err := redisCli.HMGet("key", uid).Result()
	if err != nil {
		fmt.Println("err = ", err)
	}

	for i, item := range data.([]string) {
		if item != rs[i] {
			t.Error("")
		}
	}
	return
}
