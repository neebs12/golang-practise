package main

import "fmt"

func FindMissingInt(ints []int) int {
	// extract min and max with greedy algo O(N)
	// + extract values with hash
	min := ints[0]
	max := ints[0]
	intMap := make(map[int]bool)

	for _, val := range ints {
		intMap[val] = true
		if val > max {
			max = val
		} else if val < min {
			min = val
		}
	}

	// init missingInt
	// iterate through from min to max
	// -- access the map, if any are missing, this is the missing int, we return that int

	var missingInt int

	for i := min; i <= max; i += 1 {
		if _, ok := intMap[i]; !ok {
			missingInt = i
			break
		}
	}

	return missingInt
}

func main() {
	foo := []int{2, 3, 0, 6, 1, 5}
	bar := []int{8, 2, 3, 9, 4, 7, 5, 0, 6}

	fmt.Println(FindMissingInt(foo))
	fmt.Println(FindMissingInt(bar))
}
