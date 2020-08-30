package world

import (
	"fmt"
	"time"
)

// World is the context of the world, its constants (ex: size)
type World struct {
	x            int
	y            int
	emptySymbol  string
	lifeSymbol   string
	stepDuration time.Duration
	data         *[]bool
}

// DefaultWorld is a utility function to instanciate a default world
func DefaultWorld() *World {
	result := &World{
		x:            20,
		y:            20,
		emptySymbol:  ".",
		lifeSymbol:   "o",
		stepDuration: time.Second,
		data:         &[]bool{},
	}
	size := result.x * result.y
	for i := 0; i < size; i++ {
		temp := append(*result.data, false)
		result.data = &temp
	}

	// planner
	(*result.data)[(10*result.x)+10] = true
	(*result.data)[(10*result.x)+11] = true
	(*result.data)[(10*result.x)+12] = true
	(*result.data)[(9*result.x)+12] = true
	(*result.data)[(8*result.x)+11] = true

	// periodic
	// (*result.data)[(10*result.x)+9] = true
	// (*result.data)[(10*result.x)+10] = true
	// (*result.data)[(10*result.x)+11] = true

	return result
}

// Print the world
func (w *World) Print() {
	for j := 0; j < w.y; j++ {
		offset := j * w.x
		line := []string{}
		for _, curr := range (*w.data)[offset : offset+w.y] {
			if curr {
				line = append(line, w.lifeSymbol)
			} else {
				line = append(line, w.emptySymbol)
			}
		}
		fmt.Println(line)
	}
}

// Run the world
func (w *World) Run() {
	timeTicker := time.NewTicker(w.stepDuration)
	for t := range timeTicker.C {
		fmt.Println("---------- step ", t)
		w.Print()
		w.NextStep()
	}
}
