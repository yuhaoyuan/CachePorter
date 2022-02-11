package RedisHandle

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
)

type RedisPorter interface {
	Read() error
}

type redisPorter struct {
	Client *redis.Client
	*Options
}

func NewRedisPorter(opts ...Option) RedisPorter {
	options := new(Options)
	for _, o := range opts {
		o(options)
	}

	rc := &redisPorter{
		Client:  nil,
		Options: options,
	}


	return rc
}

func (r *redisPorter) Read() error {
	switch r.Type {
	case StringKey:
		data, err := r.Client.Get(r.Key).Result()
		if err != nil && err != redis.Nil {
			return err
		}
		if data == "" || err == redis.Nil {
			// if need lock

			// if need computingFunc
			if r.NeedComputing {
				valueInterface, err := r.computingFunc()
				if err != nil {
					return err
				}

				value, ok := valueInterface.(string)
				if !ok {
					valueBytes, err := json.Marshal(value)
					if err != nil {
						return errors.New("value interface not string && cannot json Marshal")
					}
					value = string(valueBytes)
				}

				r.ReturnValue = value
				_, err = r.defaultSetFunc()
				if err != nil {
					return err
				}

			}
		}

	case HashKey:

	}

	return nil
}

func (r *RedisPorter) SetString(value string) error {
	_, err := r.client.Set(r.Key, value, r.Expire).Result()
	return err
}
