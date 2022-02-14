package RedisHandle

import (
	"github.com/go-redis/redis"
)

type RedisPorter interface {
	Read(computingFuncParams ...interface{}) error
}

type redisPorter struct {
	client *redis.Client
	Key    string
	Type   redisKeyType

	*Options
}

func NewRedisPorter(key string, keyType redisKeyType, cli *redis.Client, opts ...Option) RedisPorter {
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

func (r *redisPorter) Read(computingFuncParams ...interface{}) (err error) {
	var data interface{}

	switch r.Type {
	case StringKey:
		tempData, err := defaultGetFunc(r.client, r.Key) // get
		if err != nil && err != redis.Nil {
			return err
		}
		data = tempData

	case HashKey:
		readCmd := r.readingParam[0].(ReadCmd)
		readFunc := cmd2Func[readCmd]

		tempData, err := readFunc(r.readingParam[1:])
		if err != nil && err != redis.Nil {
			return err
		}
		data = tempData
	}

	if err == redis.Nil {
		// if need lock

		// if need computingFunc
		if r.NeedComputing {
			valueInterface, err := r.computingFunc(computingFuncParams)
			if err != nil {
				return err
			}

			r.ReturnValue = valueInterface
			err = r.cachingFunc(valueInterface) // Set
			if err != nil {
				return err
			}
		}
	}
	r.ReturnValue = data

	return nil
}
