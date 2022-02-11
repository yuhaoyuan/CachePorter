package RedisHandle

import "time"

type Options struct {
	Key     string
	LockKey string
	Type    redisKeyType
	Expire  time.Duration

	ReturnValue interface{}

	computingFunc func(param ...interface{}) (interface{}, error)
	NeedComputing bool
	readingFunc   func() (string, error)
	parsingFunc   func()
	cachingFunc   func(data interface{}) error
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{

	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func DefaultOptions() Option {
	return func(options *Options) {
		options.Key = "default_key"
		options.LockKey = ""
		options.Type = StringKey
		options.Expire = time.Second
		options.ReturnValue = ""

		options.NeedComputing = false
	}
}

func computingFuncOption(computingFunc func(param ...interface{}) (interface{}, error)) Option {
	return func(options *Options) {
		options.computingFunc = computingFunc
	}
}
