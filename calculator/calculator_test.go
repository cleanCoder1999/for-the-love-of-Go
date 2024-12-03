package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64

	want float64
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {

		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 3, want: -1},
		{a: -1, b: -1, want: 0},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {

		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: -1, b: -1, want: 1},
		{a: 5, b: 0, want: 0},
	}

	for _, tc := range testCases {

		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2500, b: 2, want: 1250},
		{a: -1, b: -1, want: 1},
		{a: 5, b: 3, want: 1.6666666},
	}

	for _, tc := range testCases {

		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if !closeEnough(tc.want, got, 0.00001) {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 144, want: 12},
		{a: 2, want: 1.41421356},
		{a: 25, want: 5},
	}

	for _, tc := range testCases {

		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if !closeEnough(tc.want, got, 0.00001) {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

// ############## INVALID INPUT tests
func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Sqrt(-1)
	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}
