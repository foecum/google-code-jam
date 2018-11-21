package main

import (
	"fmt"
	"strings"
)

func main() {
	var numTestCases int
	fmt.Scan(&numTestCases)

	for i := 1; i <= numTestCases; i++ {
		var pancakes string
		fmt.Scan(&pancakes)
		numOfFlips := flipPancakes(pancakes, i)

		fmt.Printf("Case #%d: %d\n", i, numOfFlips)
	}
}

func flipPancakes(pancakes string, index int) int {
	var numOfFlips int
	pancakesArray := strings.Split(reverse(pancakes), "")
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
	return numOfFlips
}

func reverse(s string) string {
	n := len(s)

	rsl := make([]string, n)

	for _, v := range s {
		n--
		rsl[n] = string(v)
	}
	return strings.Join(rsl, "")
}
