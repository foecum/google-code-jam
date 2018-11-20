package main

import (
	"fmt"
)

func main() {
	var numTestCases int
	fmt.Scan(&numTestCases)

	for i := 1; i <= numTestCases; i++ {
		var numOfFlips int

		var pancakes string
		fmt.Scan(&pancakes)
		pancakesArray := reverse(pancakes)
		for i, v := range pancakesArray {
			if v == "-" {
				flipPancakes(&pancakesArray, i)
				numOfFlips++
			}
		}

		fmt.Printf("Case #%d: %d\n", i, numOfFlips)
	}
}

func flipPancakes(pancakes *[]string, index int) {
	for i := index; i < len(*pancakes); i++ {
		if (*pancakes)[i] == "-" {
			(*pancakes)[i] = "+"
			continue
		}
		if (*pancakes)[i] == "+" {
			(*pancakes)[i] = "-"
		}
	}
}

func reverse(s string) []string {
	n := len(s)

	rs := make([]string, n)

	for _, v := range s {
		n--
		rs[n] = string(v)
	}
	return rs
}
