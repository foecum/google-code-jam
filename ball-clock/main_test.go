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
	}

	for _, tt := range tc {
		queue := make([]int, tt.numBalls)
		cycles := calculateCycles(&queue)
		if cycles != tt.numCycles {
			t.Errorf("expected %d but got %d", tt.numCycles, cycles)
		}
	}
}

func BenchmarkCalculateCycles(b *testing.B) {
	for n := 0; n < b.N; n++ {
		queue := make([]int, 45)
		calculateCycles(&queue)
	}
}
