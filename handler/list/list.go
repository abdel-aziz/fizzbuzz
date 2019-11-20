package list

import (
	"fizzbuzz/app"
	"fizzbuzz/encoder"

	"net/http"
)

func Register(register app.HandlerRegister) {
	register.Get("/list", ListHandler)
}

func ListHandler(ctx *app.Context, r *http.Request, w http.ResponseWriter) {
	// #1: make params, if an error occured return status 400
	params, err := NewParams(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// #2: call service to compute fizzbuzz list
	// call service fizzbuzz
	data := ctx.Fizzbuzzer.Compute(params.Int1, params.Int2, params.String1, params.String2, params.Limit)

	// #3: generate JSON output computed list
	err = encoder.Encode(w, encoder.EncoderKind_JSON, data)
	if err != nil {
		panic(err)
	}
}
