package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

// function to test Add function
func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)

	// check expression
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

// function to test Subtract function
func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)

	// check expression
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 9
	got := calculator.Multiply(3, 3)

	// check expression
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

}

func TestAddStruct(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	// Loop each testCase array
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtractStruct(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 7, b: 1, want: 6},
		{a: 5, b: 2, want: 3},
	}

	// Loop each testCase array
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiplyStruct(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 7, b: 1, want: 7},
		{a: 5, b: 2, want: 10},
	}

	// Loop each testCase array
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideStruct(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
		{a: 1, b: 3, want: 0.333333333},
	}

	// Loop each testCase array
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		// check expression for validity, exit test if invalid input detected
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}

// func TestSqrt(t *testing.T) {
// 	t.Parallel()
// 	type testCase struct {
// 		a, b float64
// 		want float64
// 	}
// 	testCases := []testCase {
// 		{a: 4, b: }
// 	}
// }

// t.Parallel >> tells Go to run test concurrently with other test
// set expected output
// call related function and pass params
// check if return value is equal to expected output
// display error message if return value is not the expected output

// := used to declare and initialize variable in concise manner
// other we would have to first declare the variable, then assign it to a value
// if := is used, we can declare and assign the variable simultaneously

// Go functions can return multiple values
// Error is a special type in Go
// Typical multiple return value is data and error value
// Pattern is called something and error

// Rule for testing >> one behaviour, one test (focused test)
// 'Something and Error' for multiple return values indicating error during execution
// Can create multiple test for one function, to test and check multiple behaviours, including valid and invalid input and executions

// 'Red, Green Refactor' process when creating a function
// Test for failures first
// Tweak for success second
// Refactor after both failures and success is checked
