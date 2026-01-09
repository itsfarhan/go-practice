// Package calculator does simple calculations.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0{
		return 0, fmt.Errorf("square root of negative number not allowed: %f",a)
	}
	return math.Sqrt(a), nil
}
