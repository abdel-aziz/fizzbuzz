package encoder

import (
	"fizzbuzz/encoder/json"

	"net/http"

	"fmt"
)

func Encode(w http.ResponseWriter, kind EncoderKind, data interface{}) error {
	var encoder Encoder
	switch kind {
	case EncoderKind_JSON:
		encoder = &json.JSONEncoder{W: w}
	default:
		return fmt.Errorf("Encoder not implemented")
	}

	return encoder.Encode(data)
}
