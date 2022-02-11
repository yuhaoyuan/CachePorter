package RedisHandle

/*

string 类型redis-key的操作

*/


func (r *RedisPorter) defaultGetFunc() (string, error) {
	return r.client.Get(r.Key).Result()
}

func (r *RedisPorter) defaultSetFunc() (string, error) {
	return r.client.Set(r.Key, r.ReturnValue, r.Expire).Result()
}

func (r *RedisPorter) defaultDelFunc() error {
	return r.client.Del(r.Key).Err()
}


