package handler

import (
	"fizzbuzz/app"
	"fizzbuzz/handler/list"
)

func RegisterRoutes(register app.HandlerRegister) {
	list.Register(register)
}
