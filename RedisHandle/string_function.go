package RedisHandle

import "github.com/go-redis/redis"

/*

string 类型redis-key的操作

*/

func defaultGetFunc(cli *redis.Client, key string) (string, error) {
	return cli.Get(key).Result()
}

func (r *redisPorter) defaultSetFunc() (string, error) {
	return r.client.Set(r.Key, r.ReturnValue, r.Expire).Result()
}

func (r *redisPorter) defaultDelFunc() error {
	return r.client.Del(r.Key).Err()
}
