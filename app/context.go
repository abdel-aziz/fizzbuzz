package app

import (
	"fizzbuzz/service/fizzbuzzer"
)

type AppContext struct {
	Fizzbuzzer fizzbuzzer.Fizzbuzzer
}

func (this *AppContext) InitContext(config Config) error {
	var err error
	this.Fizzbuzzer, err = fizzbuzzer.MakeFizzbuzzer(config.Fizzbuzzer)
	if err != nil {
		return err
	}
	return nil
}

func (this *AppContext) SetContext(context AppContext) {
	*this = context
}

type Context struct {
	AppContext
}
