package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel() // it tells Go that this test can be run in parallel with other tests; concurrency
	var want float64 = 4 // expected value
	got := calculator.Add(2, 2) // package.functionName(variables); result stored in 'got'
	if want != got { // comparing expected and actual values
		t.Errorf("want %f, got %f", want, got) // error message if they don't match
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(2, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
