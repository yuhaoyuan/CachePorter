package example

import (
	"CachePorter/RedisHandle"
	"CachePorter/RedisHandle/service"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"strconv"
	"testing"
)

func newMiniRedisClient(host string, port int) *redis.Client {
	addr := fmt.Sprintf("%v:%v", host, port)
	options := &redis.Options{
		Network: "tcp",
		Addr:    addr,
	}
	client := redis.NewClient(options)
	return client
}

func InitMiniRedisClient() *redis.Client {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	p, _ := strconv.Atoi(s.Port())
	return newMiniRedisClient(s.Host(), p)
}

func stringComputingFunc(params ...interface{}) (interface{}, error) {
	uid := params[0].(int64)
	return strconv.FormatInt(uid, 10), nil
}

func TestDefaultString(t *testing.T) {
	uid := "10001"
	key := "test-key"
	redisCli := InitMiniRedisClient()
	defer func() {
		redisCli.Del(key)
	}()

	readingParam := []interface{}{redisHandle.Get, uid}
	computingParam := []interface{}{uid}
	rPorter := service.NewRedisPorter(key, redisHandle.HashKey, redisCli, service.DefaultOptions(), service.ComputingFuncOption(computingParam, stringComputingFunc), service.ReadingFuncOption(readingParam))

	data, err := rPorter.Read()
	if err != nil {
		fmt.Println("err = ", err)
	}

	rs, err := redisCli.Get("key").Result()
	if err != nil {
		fmt.Println("err = ", err)
	}

	if data.(string) != rs {
		t.Error("data should = rs")
	}
}
