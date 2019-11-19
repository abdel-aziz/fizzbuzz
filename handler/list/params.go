package list

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	PARAMS_STRING1 = "string1"
	PARAMS_STRING2 = "string2"
	PARAMS_INT1    = "int1"
	PARAMS_INT2    = "int2"
	PARAMS_LIMIT   = "limit"
)

type Params struct {
	String1 string
	String2 string
	Int1    int
	Int2    int
	Limit   int
}

func NewParams(r *http.Request) (*Params, error) {
	int1, err := formValueToPositiveInt(r, PARAMS_INT1)
	if err != nil {
		return nil, err
	}

	int2, err := formValueToPositiveInt(r, PARAMS_INT2)
	if err != nil {
		return nil, err
	}

	limit, err := formValueToPositiveInt(r, PARAMS_LIMIT)
	if err != nil {
		return nil, err
	}

	return &Params{
		String1: r.FormValue(PARAMS_STRING1),
		String2: r.FormValue(PARAMS_STRING2),
		Int1:    int1,
		Int2:    int2,
		Limit:   limit,
	}, nil
}

func formValueToPositiveInt(r *http.Request, key string) (int, error) {
	param := r.FormValue(key)

	val, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		return 0, err
	}

	if val <= 0 {
		return 0, fmt.Errorf("Params %s is not positive value", key)
	}

	return int(val), nil
}
