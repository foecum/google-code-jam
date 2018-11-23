package main

import (
	"fmt"
	"sort"
	"time"
)

var zeroIndex = -1
var minuteQueue = make([]int, 0)
var fiveMinuteQueue = make([]int, 0)
var hourQueue = make([]int, 0)

func main() {
	var numOfBalls int
	fmt.Scan(&numOfBalls)

	ballQueue := make([]int, numOfBalls)
	startTime := time.Now()
	numOfCyles := calculateCycles(&ballQueue)
	since := time.Since(startTime)
	fmt.Printf("%d balls cycle after %d days.\n", numOfBalls, numOfCyles/2)
	fmt.Printf("Run for %s\n", since)
	fmt.Printf("Run for %.02fs\n", since.Seconds())
}

func calculateCycles(ballQueue *[]int) int {
	var cycles int
	if len(*ballQueue) >= 27 && len(*ballQueue) <= 127 {
		for i := range *ballQueue {
			(*ballQueue)[i] = i + 1
		}

		for {
			cycles += popQueue(ballQueue)
			if sort.IntsAreSorted(*ballQueue) && !contains(ballQueue, 0) {
				break
			}
		}

	}
	return cycles
}

func contains(balls *[]int, search int) bool {
	for _, v := range *balls {
		if v == search {
			return true
		}
	}
	return false
}

func popQueue(ballQueue *[]int) int {
	var zeroFound bool
	var cycle int

	ballToPop := (*ballQueue)[0]
	for i := range *ballQueue {
		if i < len(*ballQueue)-1 {
			(*ballQueue)[i] = (*ballQueue)[i+1]
			if (*ballQueue)[i] == 0 && !zeroFound {
				zeroIndex = i
				zeroFound = true
			}
		}
	}

	(*ballQueue)[len(*ballQueue)-1] = 0
	if zeroIndex == -1 {
		zeroIndex = len(*ballQueue) - 1
	}

	switch len(minuteQueue) {
	case 4:
		cycle = resetMinuteHand(ballToPop, ballQueue)
	default:
		minuteQueue = append(minuteQueue, ballToPop)
	}

	return cycle
}

func resetMinuteHand(ball int, ballQueue *[]int) int {
	n := len(minuteQueue)

	for range minuteQueue {
		n--
		(*ballQueue)[zeroIndex] = minuteQueue[n]
		zeroIndex++
	}

	minuteQueue = nil
	if len(fiveMinuteQueue) == 11 {
		return resetFiveMinHand(ball, ballQueue)
	}

	fiveMinuteQueue = append(fiveMinuteQueue, ball)

	return 0
}

func resetFiveMinHand(ball int, ballQueue *[]int) int {
	n := len(fiveMinuteQueue)

	for range fiveMinuteQueue {
		n--
		(*ballQueue)[zeroIndex] = fiveMinuteQueue[n]
		zeroIndex++
	}

	fiveMinuteQueue = nil
	if len(hourQueue) == 11 {
		resetHourHand(ball, ballQueue)
		return 1
	}
	hourQueue = append(hourQueue, ball)
	return 0
}

func resetHourHand(ball int, ballQueue *[]int) {
	n := len(hourQueue)

	for range hourQueue {
		n--
		(*ballQueue)[zeroIndex] = hourQueue[n]
		zeroIndex++
	}
	(*ballQueue)[len(*ballQueue)-1] = ball
	hourQueue = nil
}
