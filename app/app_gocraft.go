package app

import (
	"github.com/gocraft/web"

	"net/http"
)

const ROUTER_GOCRAFT = "gocraft"

type AppGoCraft struct {
	*Context
	Router *web.Router
}

func NewAppGoCraft() App {
	return &AppGoCraft{Context: &Context{}, Router: web.New(Context{})}
}

func (this *AppGoCraft) InitMiddlewares(config Config) {
	this.Router.Middleware(this.ContextMiddleware)
	this.Router.Middleware(PanicCatcherMiddleware)
}

func (this *AppGoCraft) ContextMiddleware(ctx *Context, w web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	ctx.AppContext = this.AppContext
	next(w, r)
}

func (this *AppGoCraft) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.Router.ServeHTTP(w, r)
}

func (this *AppGoCraft) Get(method string, handler func(ctx *Context, req *http.Request, w http.ResponseWriter)) {
	this.Router.Get(method, func(ctx *Context, rw web.ResponseWriter, r *web.Request) {
		handler(ctx, r.Request, rw)
	})
}

func (this *AppGoCraft) Post(method string, handler func(ctx *Context, req *http.Request, w http.ResponseWriter)) {
	this.Router.Post(method, func(ctx *Context, rw web.ResponseWriter, r *web.Request) {
		handler(ctx, r.Request, rw)
	})
}

func (this *AppGoCraft) Put(method string, handler func(ctx *Context, req *http.Request, w http.ResponseWriter)) {
	this.Router.Put(method, func(ctx *Context, rw web.ResponseWriter, r *web.Request) {
		handler(ctx, r.Request, rw)
	})
}

func (this *AppGoCraft) Delete(method string, handler func(ctx *Context, req *http.Request, w http.ResponseWriter)) {
	this.Router.Delete(method, func(ctx *Context, rw web.ResponseWriter, r *web.Request) {
		handler(ctx, r.Request, rw)
	})
}
