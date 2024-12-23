package calculator

import (
	"errors"
)

// function that takes 2 param and return the sum of params
func Add(a, b float64) float64 {
	return a + b
}

// function that takes 2 param and return the difference of params
func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	return a / b, nil
}
