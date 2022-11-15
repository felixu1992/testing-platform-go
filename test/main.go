package main

import (
	"fmt"
	"time"
)

type Option func(options *Options)

type Options struct {
	TimeOut     time.Duration
	RetryMaxNum int
}

func loadOp(option ...Option) *Options {
	options := new(Options)
	for _, e := range option {
		e(options)
	}
	return options
}

func Handler(option ...Option) {
	op := loadOp(option...)
	fmt.Printf("%v", op)
}

func main() {
	Handler()
	Handler(func(options *Options) {
		options.TimeOut = time.Millisecond
	})
	Handler(func(options *Options) {
		options.RetryMaxNum = 1
	})
	Handler(func(options *Options) {
		options.RetryMaxNum = 1
	}, func(options *Options) {
		options.TimeOut = time.Millisecond
	})
}
