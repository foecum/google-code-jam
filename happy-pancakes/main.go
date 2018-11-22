package main

import (
	"fmt"
)

func main() {
	var numTestCases int
	fmt.Scan(&numTestCases)

	for i := 1; i <= numTestCases; i++ {
		var pancakes string
		fmt.Scan(&pancakes)
		numOfFlips := flipPancakes([]rune(pancakes), i)

		fmt.Printf("Case #%d: %d\n", i, numOfFlips)
	}
}

func flipPancakes(pancakes []rune, index int) int {
	var numOfFlips int
	reverse(&pancakes)
	for i, v := range pancakes {
		if v != '-' {
			continue
		}
		for j := i; j < len(pancakes); j++ {
			switch pancakes[j] {
			case '-':
				pancakes[j] = '+'
			case '+':
				pancakes[j] = '-'
			}
		}
		numOfFlips++
	}
	return numOfFlips
}

func reverse(runes *[]rune) {
	for i, j := 0, len(*runes)-1; i < j; i, j = i+1, j-1 {
		(*runes)[i], (*runes)[j] = (*runes)[j], (*runes)[i]
	}
}
