package sum_test

import (
	"csvparserchallenge/src/sum"
	"testing"
)

func TestSum(t *testing.T) {
	number1 := 1
	number2 := 4

	result := sum.Sum(number1, number2)

	expected := 5
	if result != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, expected)
	}
}
