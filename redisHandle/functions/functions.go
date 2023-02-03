package functions

import "CachePorter/redisHandle/constant"

var (
	Cmd2GetFunc = map[constant.ReadCmd]func(params ...interface{}) (interface{}, error){
		constant.Get:     DefaultGetFunc,
		constant.HGet:    DefaultHGetFunc,
		constant.HMGet:   DefaultHMGetFunc,
		constant.HGetAll: DefaultHGetAllFunc,
	}

	Cmd2SetFunc = map[constant.ReadCmd]func(params ...interface{}) (interface{}, error){
		constant.Get:   DefaultSetFunc,
		constant.HGet:  DefaultHSet,
		constant.HMGet: DefaultHMSetFunc,
		constant.HGetAll: DefaultHMSetFunc,

	}
)
