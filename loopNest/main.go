package main

import (
	"github.com/jia-hua/algo-test/loopNest/internal/nestedLoop"
	"github.com/jia-hua/algo-test/loopNest/internal/standardLoop"
)

func main() {
	// See: https://en.wikipedia.org/wiki/Loop_nest_optimization

	problemSize := 1000

	standardLoop.Run(problemSize)
	nestedLoop.Run(problemSize)
}
