package main

import "fmt"

func TwoSum(sli []int) bool {
	myMap := make(map[int]bool)
	for _, val := range sli {
		myMap[val] = true
	} // O(N)

	for i := 0; i <= 10; i += 1 {
		lowVal := i
		highVal := 10 - i
		if lowVal == highVal { // does not account for 5, 5
			continue
		}

		_, okLow := myMap[lowVal]
		_, okHigh := myMap[highVal]
		if okLow && okHigh {
			return true
		}
	} // O(1)

	return false
}

func main() {
	myCollection := []int{2, 0, 4, 1, 7, 9}
	fmt.Println(TwoSum(myCollection))
	myCollection2 := []int{1, 2, 3, 4, 5}
	fmt.Println(TwoSum(myCollection2))
}
