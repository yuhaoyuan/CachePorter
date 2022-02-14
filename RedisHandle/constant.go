package RedisHandle

type redisKeyType uint8

const (
	StringKey redisKeyType = iota
	HashKey
	ListKey
	SetKey
	ZSetKey
)

type ReadCmd uint8

const (
	HGet  = 1
	HMGet = 2
)
