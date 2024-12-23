// Package calculator does simple calculations.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Add takes Two numbers and returns the result of adding
// them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two number a and b, and
// returns the result of multiplying  a and b.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide function
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divition by zero not allowed")
	}
	return a / b, nil
}

// Sqrt function
func Sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, fmt.Errorf("square root of negative number not allowed: %f", input)
	}

	return math.Sqrt(input), nil

}
