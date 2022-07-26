package main

import "fmt"

func UniquePaths(col, row, pathNum int, state string) int {
	// we know that each path is captured by either moving right or down
	// moving right is capured by col - 1
	// moving down is captured by row - 1
	// a path is signified by
	// -- changing directions (change in decrementing) but only not being 'forced' to change decrement if a boundary is reached.
	// to signify the end of a path, col == 1 && row == 1
	// to signify a forced change in direction (same path),
	// -- col boundary col == 1 && row > 1 (row is decremented, no change in path num)
	// -- row boundary col > 1 && row == 1 (col is decremented, no change in path num)

	// this is GCL, if row or col are 0, they are out of bounds
	if row == 0 || col == 0 {
		return 0
	}

	// therefore basecase is - reached finish line:
	if row == 1 && col == 1 {
		return pathNum
	}

	// isRight determines the previous decrement
	incrementRightPathNum := 0
	incrementDownPathNum := 0

	switch state {
	case "BEGIN":
		incrementDownPathNum = 1
		incrementRightPathNum = 1
	case "PREV_RIGHT": // ***
		// if changing dirs as not being forced (more rows exists)
		if row != 1 {
			incrementDownPathNum = 1
		}
	case "PREV_DOWN": // *** this depends!
		// if changing dirs as not being forced (more cols exists)
		if col != 1 {
			incrementRightPathNum = 1
		}
	}

	// -- if isRight (true), then the previous decrement is row - 1
	pathRight := UniquePaths(col, row-1, pathNum+incrementRightPathNum, "PREV_RIGHT")
	// -- if isRight (false), then the previous decrement is col - 1
	pathLeft := UniquePaths(col-1, row, pathNum+incrementDownPathNum, "PREV_DOWN")

	return pathRight + pathLeft
}

func main() {
	fmt.Println(UniquePaths(0, 0, 0, "BEGIN")) // 0
	fmt.Println(UniquePaths(1, 0, 0, "BEGIN")) // 0
	fmt.Println(UniquePaths(0, 1, 0, "BEGIN")) // 0
	fmt.Println(UniquePaths(1, 1, 0, "BEGIN")) // 0, yes 1, 1 is finish line
	fmt.Println(UniquePaths(1, 2, 0, "BEGIN")) // 1?
	fmt.Println(UniquePaths(2, 2, 0, "BEGIN")) // 4? (curr), 2(correct)
	fmt.Println(UniquePaths(7, 3, 0, "BEGIN")) // 4? (curr), 2(correct)

}
