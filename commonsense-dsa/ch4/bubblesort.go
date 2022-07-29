package main

import "fmt"

// sorts in place
func BubbleSort(ints []int) {
	for i := 0; i < len(ints); i += 1 {
		// is - i as last element is always at the correct place
		for j := 1; j < len(ints)-i; j += 1 {
			if ints[j-1] > ints[j] {
				// if left hand side is greater, swap places
				// -- the two pointers move with the iteration
				ints[j-1], ints[j] = ints[j], ints[j-1]
			}
		}
	}
}

func main() {
	foo := []int{11, 55, 88, 23, 457, 34, 989, 3, 878, 443}
	fmt.Println(foo)
	BubbleSort(foo)
	fmt.Println(foo)
}
