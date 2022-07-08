package main

import (
	"fmt"
)

func main() {
	arr := [2]int{11, 22}
	arr2 := arr

	fmt.Println(arr, arr2) // [11, 22] [11, 22]
	arr[0] = 1000
	fmt.Println(arr, arr2) // [1000, 22] [11, 22]
}
