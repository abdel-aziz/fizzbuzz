package encoder

type Encoder interface {
	Encode(data interface{}) error
}
