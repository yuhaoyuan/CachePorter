package RedisHandle


/*

Hash 类型redis-key的操作

*/


// Hash
func (r *RedisPorter) defaultHGetFunc(member string) (string, error) {
	data, err := r.client.HGet(r.Key, member).Result()
	if err != nil {
		return "", err
	}
	return data, err
}

func (r *RedisPorter) defaultHMGetFunc(members []string) ([]interface{}, error) {
	dataList, err := r.client.HMGet(r.Key, members...).Result()
	if err != nil {
		return dataList, err
	}
	return dataList, err
}

func (r *RedisPorter) defaultHSetFunc(member string, value string) (bool, error) {
	ok, err := r.client.HSet(r.Key, member, value).Result()
	return ok, err
}
