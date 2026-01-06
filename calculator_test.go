package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel() // it tells Go that this test can be run in parallel with other tests; concurrency
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 1, b: 2, want: 3},
		{a: -1, b: -1, want: -2},
		{a: 0, b: 0, want: 0},
		{a: 5, b: 4, want: 9},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f) : want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
	// var want float64 = 4        // expected value
	// got := calculator.Add(2, 2) // package.functionName(variables); result stored in 'got'
	// if want != got {            // comparing expected and actual values
	// 	t.Errorf("want %f, got %f", want, got) // error message if they don't match
	// }
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 5, b: 3, want: 2},
		{a: 2, b: 4, want: -2},
		{a: 0, b: 0, want: 0},
		{a: -1, b: -1, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f) : want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 3, want: 6},
		{a: -1, b: 4, want: -4},
		{a: 0, b: 5, want: 0},
		{a: -2, b: -3, want: 6},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	// Since divide can have two sensible outcomes (normal division and division by zero), therefore we will test both scenarios.
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 6, b: 3, want: 2},
		{a: 5, b: 2, want: 2.5},
		{a: 10, b: 5, want: 2}, // expecting 0 for division by zero case
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if tc.want != got {
			t.Errorf("Divide(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

