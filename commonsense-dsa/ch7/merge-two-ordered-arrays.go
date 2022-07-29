package main

import "fmt"

func Merge(ints1, ints2 []int) []int {
	resultInts := make([]int, 0, len(ints1)+len(ints2))
	ints1Ind := 0
	ints2Ind := 0

	// see if the index 'pointers' exceed the length of the input slices
	// -- if it doesnt, iterate the for loop
	for len(ints1) > ints1Ind || len(ints2) > ints2Ind {

		// can no longer iterate ints1, therefore must insert ints2 element to resultItns
		if len(ints1) == ints1Ind {
			resultInts = append(resultInts, ints2[ints2Ind])
			ints2Ind += 1

			// can no longer iterate ints2, therefore must insert ints1 element
		} else if len(ints2) == ints2Ind {
			resultInts = append(resultInts, ints1[ints1Ind])
			ints1Ind += 1

			// here, the we compare the actual integers in the input slices
			// if the ints2 value is less, we input the ints2 value and iterate its index ptr
		} else if ints1[ints1Ind] >= ints2[ints2Ind] {
			resultInts = append(resultInts, ints2[ints2Ind])
			ints2Ind += 1
		} else {
			resultInts = append(resultInts, ints1[ints1Ind])
			ints1Ind += 1
		}
	}
	return resultInts
}

func main() {
	foo := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	bar := []int{9, 10, 15, 34, 35, 36, 89, 100, 101}
	fmt.Println(Merge(foo, bar))
}
