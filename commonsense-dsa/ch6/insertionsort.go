package main

import "fmt"

func InsertionSort(ints []int) {
	for i := 1; i < len(ints); i += 1 {
		tempVal := ints[i]
		for j := i - 1; j >= 0; j -= 1 {
			scanVal := ints[j]
			if scanVal > tempVal { // valid, swap
				// this shifts to ints to the right, mimicking insertion sort
				ints[j], ints[j+1] = ints[j+1], ints[j]
			} else if scanVal <= tempVal { // not valid, break inner loop
				break
			}
		}
	}
}

func BookInsertionSort(ints []int) {
	for i := 1; i < len(ints); i += 1 {
		tempVal := ints[i]
		position := i - 1
		for position >= 0 {
			// if we encounter a value less than the temp value, we halt the loop
			if tempVal >= ints[position] {
				break
			} else {
				// shift to the right, by copying the data
				ints[position+1] = ints[position]
			}
			position -= 1
		}
		// is position + 1 here due to last -= 1 of the for loop
		ints[position+1] = tempVal
	}
}

func main() {
	foo := []int{11, 55, 88, 23, 457, 34, 989, 3, 878, 443}
	fmt.Println(foo)
	BookInsertionSort(foo)
	fmt.Println(foo)

	bar := []int{4, 2}
	fmt.Println(bar)
	BookInsertionSort(bar)
	fmt.Println(bar)
}
