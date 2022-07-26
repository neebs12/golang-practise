package main

import "fmt"

func AddUntil100(collection []int) int {
	if len(collection) == 0 {
		return 0
	}

	fmt.Println("RECURSION!")

	firstVal := collection[0]
	decrCollection := collection[1:]
	result := AddUntil100(decrCollection)
	if firstVal+result > 100 {
		return result
	} else {
		return firstVal + result
	}
}

func main() {
	fmt.Println(AddUntil100([]int{11, 22, 33, 44, 55, 66, 77}))
	fmt.Println(AddUntil100([]int{33, 8, 44, 3, 11}))
}
