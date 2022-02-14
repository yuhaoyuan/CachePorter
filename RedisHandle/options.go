package RedisHandle

import "time"

type Options struct {
	LockKey string
	Expire  time.Duration

	ReturnValue interface{}

	computingFunc  func(param ...interface{}) (interface{}, error)
	computingParam []interface{}
	NeedComputing  bool

	readingParam []interface{} // cmd, param1

	parsingFunc func()
	cachingFunc func(data interface{}) error
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
		options.LockKey = ""
		options.Expire = time.Second
		options.ReturnValue = ""

		options.NeedComputing = true

	}
}

func ComputingFuncOption(params []interface{}, computingFunc func(param ...interface{}) (interface{}, error)) Option {
	return func(options *Options) {
		options.computingParam = params
		options.computingFunc = computingFunc
	}
}

func ReadingFuncOption(params []interface{}) Option {
	return func(options *Options) {
		options.readingParam = params
	}
}
