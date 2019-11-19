package app

import (
	"net/http"
)

type App interface {
	HandlerRegister

	InitMiddlewares(config Config)

	InitContext(config Config)
	SetContext(context AppContext)
}

type HandlerRegister interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)

	Get(method string, handler func(ctx *Context, req *http.Request, w http.ResponseWriter))
	Post(method string, handler func(ctx *Context, req *http.Request, w http.ResponseWriter))
	Put(method string, handler func(ctx *Context, req *http.Request, w http.ResponseWriter))
	Delete(method string, handler func(ctx *Context, req *http.Request, w http.ResponseWriter))
}
