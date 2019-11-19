package app

import (
	"github.com/gocraft/web"

	"fmt"
	"runtime/debug"
)

func PanicCatcherMiddleware(ctx *Context, w web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf(string(debug.Stack()))

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			return
		}
	}()

	next(w, r)
}
