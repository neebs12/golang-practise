package main

import "fmt"

func main() {
	counter := 0
	foo := make([][]int, 3)
	// gives [[] [] []]

	for i := 0; i < len(foo); i += 1 {
		innerAry := make([]int, i+1)
		foo[i] = innerAry
		for j := 0; j < len(innerAry); j += 1 {
			innerAry[j] = counter
			counter += 1
		}
	}

	fmt.Println(foo)

	// make 2D, filled with values
	// [[0], [1, 2], [3, 4, 5]]

}
