package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ClockState stores the current state of the clock by recording
// how many balls are help by the Min, FiveMin, Hour hands and main ball holders
type ClockState struct {
	Min     []int
	FiveMin []int
	Hour    []int
	Main    []int
}

var clockState = ClockState{
	Min:     make([]int, 0),
	FiveMin: make([]int, 0),
	Hour:    make([]int, 0),
}

var minuteCount int

func main() {
	var numOfBalls, numOfMinutes int
	fmt.Scan(&numOfBalls)
	fmt.Scan(&numOfMinutes)
	clockState.Main = make([]int, numOfBalls)
	if len(clockState.Main) >= 27 && len(clockState.Main) <= 127 {
		for i := 0; i < numOfBalls; i++ {
			clockState.Main[i] = i + 1
		}
	}
	startTime := time.Now()
	calculateCycles(numOfMinutes)
	since := time.Since(startTime)

	b, err := json.Marshal(clockState)
	if err != nil {
		fmt.Printf("Could not marshall the clockState struct")
		return
	}

	fmt.Printf("%v", string(b))
	fmt.Printf("Run for %s\n", since)
	fmt.Printf("Run for %.02fs\n", since.Seconds())
}

func calculateCycles(numOfMinutes int) {
	for {
		if minuteCount >= numOfMinutes {
			break
		}
		popQueue()
		minuteCount++
	}
}

func popQueue() {
	ballToPop := clockState.Main[0]
	clockState.Main = clockState.Main[1:]

	switch len(clockState.Min) {
	case 4:
		resetMinuteHand(ballToPop)
	default:
		clockState.Min = append(clockState.Min, ballToPop)
	}
}

func resetMinuteHand(ball int) {
	n := len(clockState.Min)

	for range clockState.Min {
		n--
		clockState.Main = append(clockState.Main, clockState.Min[n])
	}

	clockState.Min = nil
	clockState.Min = make([]int, 0)
	if len(clockState.FiveMin) == 11 {
		resetFiveMinHand(ball)
		return
	}

	clockState.FiveMin = append(clockState.FiveMin, ball)
}

func resetFiveMinHand(ball int) {
	n := len(clockState.FiveMin)

	for range clockState.FiveMin {
		n--
		clockState.Main = append(clockState.Main, clockState.FiveMin[n])
	}

	clockState.FiveMin = nil
	clockState.FiveMin = make([]int, 0)
	if len(clockState.Hour) == 11 {
		resetHourHand(ball)
		return
	}
	clockState.Hour = append(clockState.Hour, ball)
}

func resetHourHand(ball int) {
	n := len(clockState.Hour)

	for range clockState.Hour {
		n--
		clockState.Main = append(clockState.Main, clockState.Hour[n])
	}

	clockState.Main = append(clockState.Main, ball)
	clockState.Hour = nil
	clockState.Hour = make([]int, 0)
}
