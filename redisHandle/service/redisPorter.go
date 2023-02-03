package service

import (
	"CachePorter/redisHandle/constant"
	"CachePorter/redisHandle/functions"
	"github.com/go-redis/redis"
)

type RedisPorter interface {
	Read() (interface{}, error)
}

type redisPorter struct {
	client *redis.Client
	Key    string

	*Options
}

func NewRedisPorter(key string, cli *redis.Client, opts ...Option) RedisPorter {
	options := new(Options)
	for _, o := range opts {
		o(options)
	}

	// defaultSetFunc

	rc := &redisPorter{
		client:  cli,
		Key:     key,
		Options: options,
	}

	return rc
}

// 这里更好的做法是 return固定类型的值，但目前没想到好的做法。
func (r *redisPorter) Read() (rs interface{}, err error) {
	var data interface{}

	readCmd := r.readingParam[0].(constant.ReadCmd)
	readFunc := functions.Cmd2GetFunc[readCmd]
	tempData, err := readFunc(r.readingParam[1:]) // get
	if err != nil && err != redis.Nil {
		return nil, err
	}
	data = tempData

	if err == redis.Nil {
		if r.NeedLock {
			// do lock
			// todo: fixme
		}

		// if need computingFunc
		if r.NeedComputing {
			valueInterface, err := r.computingFunc(r.computingParam)
			if err != nil {
				return nil, err
			}

			r.ReturnValue = valueInterface

			if r.cachingFunc == nil {
				r.cachingFunc = functions.Cmd2SetFunc[readCmd]
			}

			cachingParams := []interface{}{} // todo: 这里怎么设置呢？
			_, err = r.cachingFunc(cachingParams)
			if err != nil {
				return nil, err
			}
		}
	}
	r.ReturnValue = data

	return r.ReturnValue, nil
}
