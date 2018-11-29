package main

import "testing"

type testTable struct {
	values []float64
	result float64
}

var testpair = []testTable{
	{[]float64{1, 2}, 1.5},
	{[]float64{2, 2, 2, 2}, 2.0},
	{[]float64{-1, 1}, 0},
	{[]float64{3, 3, 3}, 3},
}

func TestAverage(t *testing.T) {

	for _, pair := range testpair {
		avg := Average(pair.values...)
		if avg != pair.result {
			t.Error(
				"For ", pair.values, " expected ", pair.result, " got ", avg,
			)
		}
	}
}
