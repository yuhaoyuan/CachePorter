package service

import (
	"CachePorter/redisHandle/constant"
	"github.com/go-redis/redis"
)

type RedisPorter interface {
	Read() (interface{}, error)
}

type redisPorter struct {
	client *redis.Client
	Key    string
	Type   constant.RedisKeyType

	*Options
}

func NewRedisPorter(key string, keyType constant.RedisKeyType, cli *redis.Client, opts ...Option) RedisPorter {
	options := new(Options)
	for _, o := range opts {
		o(options)
	}

	rc := &redisPorter{
		client:  cli,
		Key:     key,
		Type:    keyType,
		Options: options,
	}

	return rc
}

// 这里可以多态就好了，应该return 一个 固定的值和类型。
func (r *redisPorter) Read() (rs interface{}, err error) {
	var data interface{}

	readCmd := r.readingParam[0].(constant.ReadCmd)
	readFunc := constant.Cmd2Func[readCmd]
	tempData, err := readFunc(r.readingParam[1:]) // get
	if err != nil && err != redis.Nil {
		return nil, err
	}
	data = tempData

	if err == redis.Nil {
		// if need lock

		// if need computingFunc
		if r.NeedComputing {
			valueInterface, err := r.computingFunc(r.computingParam)
			if err != nil {
				return nil, err
			}

			r.ReturnValue = valueInterface
			err = r.cachingFunc(valueInterface) // Set
			if err != nil {
				return nil, err
			}
		}
	}
	r.ReturnValue = data

	return r.ReturnValue, nil
}

