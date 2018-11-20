package main

import (
	"fmt"
	"sort"
	"strings"
)

// Results represents the result from flipping a full single row of pancake
type Results struct {
	numberOfFlips int
	hasSolution   bool
}

func main() {
	var numberOftestCases int
	solutionResultsMap := make(map[int]Results)

	fmt.Scan(&numberOftestCases)

	for i := 1; i <= numberOftestCases; i++ {
		var pancakesRow string
		var numberToFlip int

		fmt.Scan(&pancakesRow)
		fmt.Scan(&numberToFlip)

		solution := Results{hasSolution: true}

		if numberToFlip <= len(pancakesRow) {
			pancakesRowArray := strings.Split(pancakesRow, "")
			for index := 0; index < len(pancakesRowArray); index++ {
				if pancakesRowArray[index] != "+" {
					if index+(numberToFlip-1) <= len(pancakesRowArray)-1 {
						solution.numberOfFlips++
						for j := 0; j < numberToFlip; j++ {
							switch pancakesRowArray[index+j] {
							case "-":
								pancakesRowArray[index+j] = "+"
							case "+":
								pancakesRowArray[index+j] = "-"
							}
						}
					}
				}
			}
			solution.hasSolution = contains(pancakesRowArray, "-")
			solutionResultsMap[i] = solution
		}
	}

	var keys []int
	for k := range solutionResultsMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		printTestCaseResult(k, solutionResultsMap[k])
	}
}

func contains(pancakes []string, pancake string) bool {
	for j := range pancakes {
		if pancakes[j] == pancake {
			return true
		}
	}
	return false
}

func printTestCaseResult(testCaseNumber int, solutionResult Results) {
	if solutionResult.hasSolution {
		fmt.Printf("Case #%d: %d\n", testCaseNumber, solutionResult.numberOfFlips)
		return
	}
	fmt.Printf("Case #%d: IMPOSSIBLE\n", testCaseNumber)
}
