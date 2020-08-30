package main

import (
	"fmt"

	"github.com/jia-hua/algo-test/asciiLife/internal/world"
)

func main() {
	fmt.Println("-------- begin world")

	myWorld := world.DefaultWorld()
	myWorld.Run()
}
