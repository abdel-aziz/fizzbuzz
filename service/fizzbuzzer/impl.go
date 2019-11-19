package fizzbuzzer

// Fizzbuzzer implementation based on basic loop for
type BasicFizzbuzzer struct{}

func NewBasicFizzbuzzer(config *BasicConfig) Fizzbuzzer {
	return &BasicFizzbuzzer{}
}

func (this *BasicFizzbuzzer) Compute(int1, int2 int, string1, string2 string, limit int) []interface{} {
	output := make([]interface{}, limit)
	for i := 1; i <= len(output); i++ {
		index := i - 1
		switch {
		case i%(int1*int2) == 0:
			output[index] = string1 + string2
		case i%int1 == 0:
			output[index] = string1
		case i%int2 == 0:
			output[index] = string2
		default:
			output[index] = i
		}
	}
	return output
}
