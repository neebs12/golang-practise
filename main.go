package main

import "fmt"

func getIntSum(nums ...int) int {
	// nums is a collection array/slice, does not matter?
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func main() {
	foo := []int{11, 11, 22, 33, 44}
	fmt.Println(getIntSum(foo...)) // <--- yes, 121
	// <--- ok slices

	bar := [5]int{11, 11, 22, 33, 44}
	fmt.Println(getIntSum(bar...)) // <--- no
}
