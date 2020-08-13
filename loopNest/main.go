package main

import (
	"fmt"
	"time"

	"github.com/jia-hua/algo-test/loopNest/internal/nestedLoop"
	"github.com/jia-hua/algo-test/loopNest/internal/standardLoop"
)

func main() {
	// See: https://en.wikipedia.org/wiki/Loop_nest_optimization

	problemSize := 1000

	standarDuration := standardLoop.Run(problemSize)
	nestedDuration := nestedLoop.Run(problemSize)

	fmt.Println("nested is ", computeEcart(nestedDuration, standarDuration), " faster in ms")
}

func computeEcart(time1, time2 time.Duration) time.Duration {
	gap := time2 - time1
	return gap
}
