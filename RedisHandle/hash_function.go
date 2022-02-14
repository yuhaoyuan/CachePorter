package RedisHandle

import "github.com/go-redis/redis"

/*

Hash 类型redis-key的操作

*/

var (
	cmd2Func = map[ReadCmd]func(params []interface{}) ([]string, error){
		HGet:  defaultHGetFunc,
		HMGet: defaultHMGetFunc,
	}
)

// Hash
func defaultHGetFunc(params []interface{}) ([]string, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	member := params[2].(string)

	data, err := cli.HGet(key, member).Result()
	if err != nil {
		return nil, err
	}
	return []string{data}, err
}

func defaultHMGetFunc(params []interface{}) ([]string, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	members := params[2].([]string)

	dataList, err := cli.HMGet(key, members...).Result()
	if err != nil {
		return dataList, err
	}
	return dataList, err
}

func (r *redisPorter) defaultHSetFunc(member string, value string) (bool, error) {
	ok, err := r.client.HSet(r.Key, member, value).Result()
	return ok, err
}
