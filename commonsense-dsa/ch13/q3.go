package main

import (
	"fmt"
	"sort"
)

var _ = sort.Ints

func FindMaxON2(collection []int) int {
	// bubble sort, then find max
	collection = append([]int{}, collection...) // copy

	for i := 0; i < len(collection); i += 1 {
		iVal := collection[i]
		for j := i + 1; j < len(collection); j += 1 {
			jVal := collection[j]
			if iVal > jVal {
				// swap
				collection[i], collection[j] = collection[j], collection[i]
				iVal = collection[j] // reassg iVal
			}
		}
	}

	return collection[len(collection)-1]
}

func FindMaxON(collection []int) int {
	currVal := collection[0]
	for _, val := range collection {
		if val > currVal {
			currVal = val
		}
	}
	return currVal
}

func FindMaxONLOGN(collection []int) int {
	collection = append([]int{}, collection...)
	sort.Ints(collection)
	return collection[len(collection)-1]
}

func main() {
	myCollection := []int{9, 11, 444, 55, 33, 77, 55, 22}
	fmt.Println(myCollection)
	fmt.Println(FindMaxON2(myCollection))
	fmt.Println(myCollection)
	fmt.Println(FindMaxON(myCollection))
	fmt.Println(myCollection)
	fmt.Println(FindMaxONLOGN(myCollection))
}
