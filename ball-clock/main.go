package main

import (
	"fmt"
	"time"
)

/**
Optimising app
-Remove slices
-Remove unneccessary allocations
-Reduce function calls
*/

var minuteQueue [4]rune
var fiveMinuteQueue [11]rune
var hourQueue [11]rune

var minuteIndex int
var fiveMinuteIndex int
var hourIndex int

var firstZeroIndex = -1

func main() {
	var numOfBalls = 123

	startTime := time.Now()
	numOfCycles := calculateCycles(numOfBalls)
	since := time.Since(startTime)
	fmt.Printf("%d balls cycle after %d days.\n", numOfBalls, numOfCycles/2)
	fmt.Printf("Run for %s\n", since)
	fmt.Printf("Run for %.02fs\n", since.Seconds())
}

func calculateCycles(numOfBalls int) int {
	ballQueue := make([]rune, numOfBalls)
	for i := range ballQueue {
		ballQueue[i] = rune(i + 1)
	}

	var numOfCyles int
	for {
		ballToPop := ballQueue[0]
		ballQueue = append(ballQueue[1:], 0)
		// for i := 0; i < len(ballQueue)-2; i += 2 {
		// 	ballQueue[i], ballQueue[i+1] = ballQueue[i+1], ballQueue[i+2]
		// }
		// ballQueue[len(ballQueue)-1] = 0

		switch firstZeroIndex < 0 {
		case true:
			firstZeroIndex = len(ballQueue) - 1
		default:
			firstZeroIndex--
		}

		switch minuteIndex < 4 {
		case true:
			minuteQueue[minuteIndex] = ballToPop
			minuteIndex++
		default:
			// minute hand is full
			minuteIndex--
			ballQueue[firstZeroIndex] = minuteQueue[minuteIndex]
			firstZeroIndex++
			minuteIndex--
			ballQueue[firstZeroIndex] = minuteQueue[minuteIndex]
			firstZeroIndex++
			minuteIndex--
			ballQueue[firstZeroIndex] = minuteQueue[minuteIndex]
			firstZeroIndex++
			minuteIndex--
			ballQueue[firstZeroIndex] = minuteQueue[minuteIndex]
			firstZeroIndex++

			// fiveMinute hand is not full
			if fiveMinuteIndex < 11 {
				fiveMinuteQueue[fiveMinuteIndex] = ballToPop
				fiveMinuteIndex++
				break
			}

			// fiveMinute hand is full
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++
			fiveMinuteIndex--
			ballQueue[firstZeroIndex] = fiveMinuteQueue[fiveMinuteIndex]
			firstZeroIndex++

			// hour hand is not full
			if hourIndex < 11 {
				hourQueue[hourIndex] = ballToPop
				hourIndex++
				break
			}
			// hour hand is full
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++
			hourIndex--
			ballQueue[firstZeroIndex] = hourQueue[hourIndex]
			firstZeroIndex++

			ballQueue[firstZeroIndex] = ballToPop
			firstZeroIndex++

			firstZeroIndex = -1
			numOfCyles++
		}

		if isSorted(ballQueue) {
			break
		}
	}

	return numOfCyles
}

func isSorted(r []rune) bool {
	n := len(r)
	for i := n - 1; i > 0; i-- {
		if r[i] < r[i-1] {
			return false
		}
	}
	return true
}
