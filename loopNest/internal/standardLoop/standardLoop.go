package standardLoop

import (
	"fmt"
	"time"
)

// Run a standard loop, and see how much time it takes to do the job
func Run(size int) time.Duration {
	start := time.Now()
	fmt.Println("standard loop")

	bigSequential(size)

	elapsed := time.Now().Sub(start)
	fmt.Println("------ standard duration = ", elapsed)

	return elapsed
}

func bigSequential(size int) {

	i, j := 0, 0
	a := make([][]int, size)
	for i := range a {
		a[i] = make([]int, size)
	}

	b := make([]int, size)

	c := make([]int, size)

	n := size

	for i = 0; i < n; i++ {
		c[i] = 0
		for j = 0; j < n; j++ {
			c[i] = c[i] + a[i][j]*b[j]
		}
	}
}
