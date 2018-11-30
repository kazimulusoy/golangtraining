package word

import (
	"testing"
)

func TestCount(t *testing.T) {
	word := "Hello World"
	s := Count(word)
	expectedLength := 2

	if s != expectedLength {
		t.Error("Length expected:" + string(expectedLength) + " got " + string(s))
	}
}

func BenchmarkCount(t *testing.B) {
	word := "Hello World"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		Count(word)
	}
}
