package app

import (
	"fizzbuzz/service/fizzbuzzer"
)

type AppContext struct {
	Fizzbuzzer fizzbuzzer.Fizzbuzzer
}

func (this *AppContext) InitContext(config Config) {
	this.Fizzbuzzer = fizzbuzzer.NewBasicFizzbuzzer()
}

func (this *AppContext) SetContext(context AppContext) {
	*this = context
}

type Context struct {
	AppContext
}
