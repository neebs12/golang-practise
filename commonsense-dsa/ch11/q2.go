package main

import "fmt"

func ExtractEvenNumbers(nums []int) []int {

	if len(nums) == 0 {
		return make([]int, 0)
	}

	if firstVal := nums[0]; firstVal%2 == 0 {
		// is even,
		return append([]int{firstVal}, ExtractEvenNumbers(nums[1:])...)
	} else {
		// is odd, ignore firstVal
		return ExtractEvenNumbers(nums[1:])
	}

}

func main() {
	intCollection := []int{11, 22, 33, 44, 55, 66, 0, 123, 122, 77, 99, 2, 3, 4, 5}
	fmt.Println(ExtractEvenNumbers(intCollection))
}
