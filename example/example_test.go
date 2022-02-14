package example

import (
	"CachePorter/RedisHandle"
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

// NewRedisClient new redis client
func NewRedisClient() *redis.Client {
	options := &redis.Options{
		Network:      "tcp",
		Addr:         "proxy.cache-codis-sg2.i.test.sz.shopee.io:8147",
		DB:           0,
		Password:     "",
		MaxRetries:   2,
		PoolSize:     2,
		DialTimeout:  time.Duration(100) * time.Millisecond,
		ReadTimeout:  time.Duration(100) * time.Millisecond,
		WriteTimeout: time.Duration(100) * time.Millisecond,
		PoolTimeout:  time.Duration(3000) * time.Millisecond,
		IdleTimeout:  time.Duration(30000) * time.Millisecond,
		Dialer:       nil,
	}
	cli := redis.NewClient(options)
	cmd := cli.Ping()
	val, err := cmd.Result()
	if err != nil || val == "" {
		panic("ping redis fail")
	}
	return cli
}

func stringComputingFunc(param ...interface{}) (interface{}, error) {
	return trueStringComputingFunc(param[0].(string))
}

func trueStringComputingFunc(uid string) (string, error) {
	return uid, nil
}

func TestDefaultString(t *testing.T) {
	uid := "10001"
	key := "test-key"
	redisCli := NewRedisClient()
	defer func() {
		redisCli.Del(key)
	}()

	rPorter := RedisHandle.NewRedisPorter(key, RedisHandle.StringKey, redisCli, RedisHandle.DefaultOptions(), RedisHandle.ComputingFuncOption([]interface{}{uid}, stringComputingFunc))

	err := rPorter.Read()
	if err != nil {
		fmt.Println("err = ", err)
	}

	rs, err := redisCli.Get("key").Result()
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("rs = ", rs)
}
