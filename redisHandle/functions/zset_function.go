package functions

import (
	"github.com/go-redis/redis"
)

// ZSet
func DefaultZRangeFunc(params []interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	l := params[2].(int64)
	r := params[3].(int64)

	data, err := cli.ZRange(key, l, r).Result()
	if err != nil {
		return nil, err
	}
	return data, err
}

func DefaultZRangeByScore(params []interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	min := params[2].(string)
	max := params[3].(string)
	offset := params[4].(int64)
	count := params[5].(int64)

	data, err := cli.ZRangeByScore(key, redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DefaultZRangeByScoreWithScore(params []interface{}) (interface{}, error) {
	cli := params[0].(*redis.Client)
	key := params[1].(string)
	min := params[2].(string)
	max := params[3].(string)
	offset := params[4].(int64)
	count := params[5].(int64)

	data, err := cli.ZRangeByScoreWithScores(key, redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}
