package main

import (
	"fmt"
)

/*
def unique_paths(rows,columns)
  return 1 if rows== 1 || columns== 1
	return unique_paths(rows- 1, columns) + unique_paths(rows,columns- 1)
end

*/

var counter int

func UniquePathsMemo(rows, cols int, memo map[string]int) int {

	// encoding rows and cols to a valid memo key
	encode := fmt.Sprintf("%d %d", rows, cols)

	if val, ok := memo[encode]; ok {
		return val
	}

	counter += 1
	// fmt.Println(counter)

	if rows == 1 || cols == 1 {
		memo[encode] = 1
		return 1
	}

	// else, we store the various values given encodings
	memo[encode] = UniquePathsMemo(rows-1, cols, memo) + UniquePathsMemo(rows, cols-1, memo)

	return memo[encode]
}

func UniquePaths(rows, cols int) int {
	if rows == 1 || cols == 1 {
		return 1
	}

	counter += 1
	// fmt.Println(counter)

	return UniquePaths(rows-1, cols) + UniquePaths(rows, cols-1)
}

func main() {
	counter = 0
	fmt.Println(UniquePaths(10, 11), "calls: ", counter)
	// 92377 calls
	counter = 0
	fmt.Println(UniquePathsMemo(10, 11, make(map[string]int)), "calls: ", counter)
	// 109 calls
}
