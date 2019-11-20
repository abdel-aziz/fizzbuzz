package fizzbuzzer

import (
	"github.com/stretchr/testify/mock"
)

// Fizzbuzzer implementation based on mock system for test purpose
type MockFizzbuzzer struct {
	mock.Mock
}

func (this *MockFizzbuzzer) Compute(int1, int2 int, string1, string2 string, limit int) []interface{} {
	args := this.Called(int1, int2, string1, string2, limit)
	var output []interface{}
	if args.Get(0) != nil {
		output = args.Get(0).([]interface{})
	}
	return output
}
