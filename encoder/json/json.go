package json

import (
	"net/http"

	"encoding/json"
)

type JSONEncoder struct {
	W http.ResponseWriter
}

func (this *JSONEncoder) Encode(data interface{}) error {
	this.W.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(this.W).Encode(data)
}
