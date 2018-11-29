package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(10)
	expected := 70
	if y != expected {
		t.Error("Expected " + string(expected) + " got " + string(y))
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(10)
	expected := 70
	if y != expected {
		t.Error("Expected " + string(expected) + " got " + string(y))
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(t *testing.B) {
	for i := 0; i < t.N; i++ {
		YearsTwo(10)
	}
}
