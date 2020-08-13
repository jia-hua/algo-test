package nestedLoop

import (
	"fmt"
	"time"
)

// Run a nested loop, and see how much time it takes to do the job
func Run(size int) {
	start := time.Now()
	fmt.Println("nested loop")

	nestedSequential(size)

	elapsed := time.Now().Sub(start)
	fmt.Println("------ nested duration = ", elapsed)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func nestedSequential(size int) {

	loopSize := 2

	i, j, x, y := 0, 0, 0, 0

	a := [100][100]uint{}
	b, c := [100]uint{}, [100]uint{}
	n := 100
	for i = 0; i < n; i += loopSize {
		c[i] = 0
		c[i+1] = 0
		for j = 0; j < n; j += loopSize {
			for x = i; x < min(i+loopSize, n); x++ {
				for y = j; y < min(j+2, n); y++ {
					c[x] = c[x] + a[x][y]*b[y]
				}
			}
		}
	}
}
