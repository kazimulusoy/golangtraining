package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	var sum int

	sum = SumTwoNumbers(1, 2)
	if sum != 3 {
		t.Error("Expected 3 but got ", sum)
	}
}
func TestSumIsWrong(t *testing.T) {
	var sum int

	sum = SumTwoNumbers(2, 2)
	if sum != 3 {
		t.Error("Expected 3 but got ", sum)
	}
}
