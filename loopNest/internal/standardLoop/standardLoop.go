package standardLoop

import (
	"fmt"
	"time"
)

// Run a standard loop, and see how much time it takes to do the job
func Run(size int) {
	start := time.Now()
	fmt.Println("standard loop")

	bigSequential(size)

	elapsed := time.Now().Sub(start)
	fmt.Println("------ standard duration = ", elapsed)
}

func bigSequential(size int) {

	i := 0
	j := 0
	a := [100][100]uint{}
	b := [100]uint{}
	c := [100]uint{}
	n := 100
	for i = 0; i < n; i++ {
		c[i] = 0
		for j = 0; j < n; j++ {
			c[i] = c[i] + a[i][j]*b[j]
		}
	}
}
