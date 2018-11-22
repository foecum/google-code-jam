package main

import "testing"

func TestReverse(t *testing.T) {
	tc := []struct {
		text     string
		expected []rune
	}{
		{
			text:     "---++",
			expected: []rune("++---"),
		},
		{
			text:     "-----",
			expected: []rune("-----"),
		},
		{
			text:     "",
			expected: []rune(""),
		},
	}
	for _, tt := range tc {
		r := []rune(tt.text)
		reverse(&r)
		if string(tt.expected) != string(r) {
			t.Errorf("test failed. expected %s but got %s\n", string(tt.expected), string(r))
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		r := []rune("---++")
		reverse(&r)
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
		var r []rune
		r = []rune(tt.text)
		result := flipPancakes(r, 1)
		if tt.expected != result {
			t.Errorf("test failed. expected %d but got %d\n", tt.expected, result)
		}
	}
}

func BenchmarkFlipPancakes(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		var r []rune
		r = []rune("---++")
		flipPancakes(r, 1)
	}
}
