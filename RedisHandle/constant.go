package RedisHandle

type redisKeyType uint8

const (
	StringKey redisKeyType = iota
	HashKey
	ListKey
	SetKey
	ZSetKey
)
