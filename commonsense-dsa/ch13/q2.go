package main

import (
	"fmt"
	"sort"
)

func FindMissingNumber(collection []int) int {
	sort.Ints(collection)
	minValue := collection[0]
	// maxValue := collection[len(collection) - 1]
	for i := 0; i < len(collection); i += 1 {
		// if the seq is broken, IS the missing number!
		if (minValue + i) != collection[i] {
			return minValue + i
		}
	}
	return 0
}

func main() {
	myInts := []int{1, 4, 7, 8, 9, 5, 3, 6}
	fmt.Println(FindMissingNumber(myInts))
}
