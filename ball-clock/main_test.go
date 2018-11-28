package main

import "testing"

func TestCalculateCycles(t *testing.T) {
	tc := []struct {
		numBalls  int
		numCycles int
	}{
		{
			numBalls:  30,
			numCycles: 30,
		},
		{
			numBalls:  45,
			numCycles: 756,
		},
		{
			numBalls:  123,
			numCycles: 217710,
		},
	}

	for _, tt := range tc {
		cycles := calculateCycles(tt.numBalls)
		if cycles != tt.numCycles {
			t.Errorf("expected %d but got %d", tt.numCycles, cycles)
		}
	}
}

func BenchmarkCalculateCycles(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calculateCycles(123)
	}
}
