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
		flipPancakes(pancakes, i)
	}
}

func flipPancakes(pancakes string, index int) {
	var numOfFlips int
	pancakesArray := reverse(pancakes)
	for i, v := range pancakesArray {
		if v != "-" {
			continue
		}
		for j := i; j < len(pancakesArray); j++ {
			switch pancakesArray[j] {
			case "-":
				pancakesArray[j] = "+"
			case "+":
				pancakesArray[j] = "-"
			}
		}
		numOfFlips++
	}
	fmt.Printf("Case #%d: %d\n", index, numOfFlips)
}

func reverse(s string) []string {
	n := len(s)

	rsl := make([]string, n)

	for _, v := range s {
		n--
		rsl[n] = string(v)
	}
	return rsl
}
