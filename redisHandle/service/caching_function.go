package service

func (r *redisPorter) defaultSetFunc() (string, error) {
	return r.client.Set(r.Key, r.ReturnValue, r.Expire).Result()
}

func (r *redisPorter) defaultDelFunc() error {
	return r.client.Del(r.Key).Err()
}

func (r *redisPorter) defaultHSetFunc(member string, value string) (bool, error) {
	ok, err := r.client.HSet(r.Key, member, value).Result()
	return ok, err
}
