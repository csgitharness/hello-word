package mathematics

import "errors"

type value struct {
	Val1 float32 `json:"val_1,omitempty"`
	Val2 float32 `json:"val_2,omitempty"`
}

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}

	return a / b, nil
}

func add(a, b float32) (float32, error) {
	return a + b, nil
}

func subtraction(a, b float32) (float32, error) {
	return a - b, nil
}

func multiply(a, b float32) (float32, error) {
	return a * b, nil
}
