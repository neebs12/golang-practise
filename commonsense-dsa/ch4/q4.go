package main

import "fmt"

// create a O(N) find max int from a collection
func findMax(collection []int) int {
	max := collection[0]
	for _, val := range collection[1:] {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	fmt.Println(findMax([]int{11, 22, 33, 900, 44, 100, 12, 666, 777}))
}
