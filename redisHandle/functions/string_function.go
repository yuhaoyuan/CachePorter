package functions

import (
	"CachePorter/redisHandle/service"
	"github.com/go-redis/redis"
)

/*

string 类型redis-key的操作

*/

func DefaultGetFunc(params []interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)

	return cli.Get(key).Result()
}

