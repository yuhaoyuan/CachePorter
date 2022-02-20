package constant

import "CachePorter/redisHandle/functions"

type RedisKeyType uint8

const (
	StringKey RedisKeyType = iota
	HashKey
	ListKey
	SetKey
	ZSetKey
)

type ReadCmd uint8

const (
	UnknownCmd ReadCmd = iota
	Get                // return string
	HGet               // return string
	HMGet              // return []string
	HGetAll            // return map[string][string]
)

var (
	Cmd2Func = map[ReadCmd]func(params []interface{}) (interface{}, error){
		Get:     functions.DefaultGetFunc,
		HGet:    functions.DefaultHGetFunc,
		HMGet:   functions.DefaultHMGetFunc,
		HGetAll: functions.DefaultHGetAllFunc,
	}
)
