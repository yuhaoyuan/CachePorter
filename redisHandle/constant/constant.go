package constant

type RedisKeyType uint8

const (
	UnknownKey RedisKeyType = iota
	StringKey
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
