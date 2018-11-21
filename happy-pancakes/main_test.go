package main

import "testing"

func TestReverse(t *testing.T) {
	tc := []struct {
		text     string
		expected string
	}{
		{
			text:     "---++",
			expected: "++---",
		},
		{
			text:     "-----",
			expected: "-----",
		},
		{
			text:     "",
			expected: "",
		},
	}
	for _, tt := range tc {
		result := reverse(tt.text)
		if tt.expected != result {
			t.Errorf("test failed. expected %s but got %s\n", tt.expected, result)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		reverse("---++")
	}
}

func TestFlipPancakes(t *testing.T) {
	tc := []struct {
		text     string
		expected int
	}{
		{
			text:     "-",
			expected: 1,
		},
		{
			text:     "-+",
			expected: 1,
		},
		{
			text:     "+-",
			expected: 2,
		},
		{
			text:     "++-++++---",
			expected: 4,
		},
	}

	for _, tt := range tc {
		result := flipPancakes(tt.text, 1)
		if tt.expected != result {
			t.Errorf("test failed. expected %d but got %d\n", tt.expected, result)
		}
	}
}

func BenchmarkFlipPancakes(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		flipPancakes("---++", 1)
	}
}
