package main

import (
	"github.com/jia-hua/algo-test/loopNest/internal/nestedLoop"
	"github.com/jia-hua/algo-test/loopNest/internal/standardLoop"
)

func main() {

	problemSize := 1000000000

	standardLoop.Run(problemSize)
	nestedLoop.Run(problemSize)
}
