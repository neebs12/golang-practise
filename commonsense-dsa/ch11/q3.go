package main

import "fmt"

var _ = fmt.Println // debug

func GetTriangularNumberByNth(nth int, counter int) int {
	// base case
	if nth == 1 {
		return 1
	}

	counter += 1

	return counter + GetTriangularNumberByNth(nth-1, counter)
}

func main() {
	fmt.Println(GetTriangularNumberByNth(1, 1))
	fmt.Println(GetTriangularNumberByNth(2, 1))
	fmt.Println(GetTriangularNumberByNth(3, 1))
	fmt.Println(GetTriangularNumberByNth(7, 1))
}
