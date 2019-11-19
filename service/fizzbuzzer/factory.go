package fizzbuzzer

import (
	"fmt"
)

func MakeFizzbuzzer(config Config) (Fizzbuzzer, error) {
	var fizzbuzzer Fizzbuzzer
	switch config.Backend {
	case "basic":
		fizzbuzzer = NewBasicFizzbuzzer(config.Basic)
	case "routine":
		fizzbuzzer = NewRoutineFizzbuzzer(config.Routine)
	case "split":
		fizzbuzzer = NewSplitFizzbuzzer(config.Split)
	default:
		return nil, fmt.Errorf("MakeFizzbuzzer: Unknown backend %s", config.Backend)
	}
	return fizzbuzzer, nil
}
