package functions

import (
	"github.com/go-redis/redis"
	"time"
)

/*

string 类型redis-key的操作

*/

func DefaultGetFunc(params ...interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)

	return cli.Get(key).Result()
}

func DefaultSetFunc(params ...interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	value := params[2].(string)
	expiration := params[3].(time.Duration)

	return cli.Set(key, value, expiration).Result()
}
