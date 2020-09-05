package world

// NextStep compute next state of the world
func (w *World) NextStep() {
	futureData := make([]bool, len(w.data))
	copy(futureData, w.data)

	for j := 1; j < w.y-1; j++ {
		for i := 1; i < w.x-1; i++ {
			subGrid := []bool{}
			x := j*w.x + i
			subGrid = append(subGrid, w.data[x-w.x-1:x-w.x+2]...)
			subGrid = append(subGrid, w.data[x-1:x+2]...)
			subGrid = append(subGrid, w.data[x+w.x-1:x+w.x+2]...)

			cellNextState := NextCellStep(subGrid)
			futureData[x] = cellNextState
		}
	}

	w.data = futureData
}

// NextCellStep compute the next state of the center of a subGrid (3x3)
// true = alive, false = dead
func NextCellStep(subGrid []bool) bool {
	currentCellAlive := subGrid[4]
	neighbors := append(subGrid[0:4], subGrid[5:]...)
	livingNeighbors := countLivingCell(neighbors)

	if currentCellAlive {
		// living = 2-3 living neigbor
		if livingNeighbors >= 2 && livingNeighbors <= 3 {
			return true
		}
	} else {
		// born = 3 living neighbor
		if livingNeighbors == 3 {
			return true
		}
	}

	// death = default
	return false
}

func countLivingCell(cellList []bool) int {
	res := 0
	for _, isAlive := range cellList {
		if isAlive {
			res++
		}
	}

	return res
}
