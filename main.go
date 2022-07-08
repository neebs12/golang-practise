package main

import (
	"fmt"
)

func main() {
	arr := [6]int{9, 11, 6, 42, 23, 600}
	sli := arr[2:5]
	fmt.Printf("An array is of type: %T, a slice is of type: %T\n", arr, sli)
	// <--- [6]int is an array, []int is a slice
	fmt.Println(arr, sli)
	sli2 := arr[:]
	// arr[len(arr) - 1] = 599 // <-- affects all refs
	fmt.Println(arr, sli2)
	sli3 := sli[1:2]
	fmt.Println(sli3)
	sli4 := sli[:cap(sli)] // expands to parent array
	fmt.Println(sli4)

	var sli5 []int = []int{11, 22, 33, 44, 55}
	fmt.Println(sli5)

	fmt.Println("-----------------")
	sli6 := append(sli, 100)
	// <--- appending value MUST be of same type as elements of the appending array
	// <--- think, only way that a collection can be accepted is if the first arg is a 2D collection or higher
	fmt.Println(sli6)

	sli7 := make([]int, 5)
	fmt.Println(sli7)
}
