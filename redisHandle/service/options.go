package service

import "time"

type Options struct {
	LockKey       string
	NeedLock      bool
	NeedComputing bool
	Expire        time.Duration
	ReturnValue   interface{}

	computingFunc  func(param ...interface{}) (interface{}, error)
	computingParam []interface{}
	readingParam   []interface{} // cmd, param1
	parsingFunc    func()
	cachingFunc    func(param ...interface{}) (interface{}, error)
}

type Option func(*Options)

func NewOptions(opts ...Option) Options {
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
		if computingFunc != nil {
			options.computingFunc = computingFunc
		}
	}
}

func ReadingFuncOption(params []interface{}) Option {
	return func(options *Options) {
		options.readingParam = params
	}
}

func CachingFuncOption(cachingFunc func(param ...interface{}) (interface{}, error)) Option {
	return func(options *Options) {
		if cachingFunc != nil {
			options.cachingFunc = cachingFunc
		}
	}
}
