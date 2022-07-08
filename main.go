package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	// arr = {11, 22, 33} // <-- no
	// arr = [11, 22, 33] // <-- no
	arr[0] = 11 // yes
	arr[1] = 22
	arr[2] = 33 
	fmt.Println(arr, len(arr))

	sum := 0
	for i := 0; i < len(arr); i += 1{
		sum += arr[i]
	}
	fmt.Println(sum)

	arr2D := [3][2]int{{11, 22}, {33, 44}, {55, 66}}
	fmt.Println(arr2D, len(arr2D))

	sum = 0 // reassg
	for i := 0 ; i < len(arr2D); i += 1{
		for j := 0; j < len(arr2D[i]); j += 1{
			sum += arr2D[i][j]
		}
	}
	fmt.Println(sum)
}
