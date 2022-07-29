package main

import "fmt"

// sorts in place
func SelectionSort(ints []int) {
	for i := 0; i < len(ints); i += 1 {
		quasiMin := ints[i]
		currMin := ints[i]
		indToSwap := i // otherwise, will perform swap on itself (which is a harmless operation)
		for j := i + 1; j < len(ints); j += 1 {
			currVal := ints[j]
			if quasiMin > currVal && currMin > currVal {
				currMin = currVal
				indToSwap = j
			}
		}

		// here, we perform the swap
		ints[i], ints[indToSwap] = ints[indToSwap], ints[i]
	}
}

func main() {
	foo := []int{11, 55, 88, 23, 457, 34, 989, 3, 878, 443}
	fmt.Println(foo)
	SelectionSort(foo)
	fmt.Println(foo)
}
