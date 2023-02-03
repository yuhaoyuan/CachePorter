package functions

import (
	"github.com/go-redis/redis"
)

/*

Hash 类型redis-key的操作

*/

// Hash
func DefaultHGetFunc(params ...interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	member := params[2].(string)

	data, err := cli.HGet(key, member).Result()
	if err != nil {
		return nil, err
	}
	return data, err
}

func DefaultHMGetFunc(params ...interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	members := params[2].([]string)
	result := []string{}

	dataInterfaceList, err := cli.HMGet(key, members...).Result()
	if err != nil {
		return result, err
	}
	for _, item := range dataInterfaceList {
		result = append(result, item.(string))
	}
	return result, err
}

func DefaultHGetAllFunc(params ...interface{}) (result interface{}, err error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)

	result, err = cli.HGetAll(key).Result()
	if err != nil {
		return
	}
	return
}

func DefaultHSet(params ...interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	member := params[2].(string)
	v := params[3].([]interface{})

	return cli.HSet(key, member, v).Result()
}

func DefaultHMSetFunc(params ...interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	value := params[2].(map[string]interface{})

	return cli.HMSet(key, value).Result()
}
