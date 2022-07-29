package main

import "fmt"

func BinarySearch(searchVal int, ints []int) bool {

	lowerBound := 0
	upperBound := len(ints) - 1

	for lowerBound <= upperBound {
		midPoint := (upperBound + lowerBound) / 2
		midPointValue := ints[midPoint]

		if midPointValue < searchVal {
			// need to look to the left of the midPoint
			// -- -1 for the shift to the left
			lowerBound = midPoint + 1
		} else if midPointValue > searchVal {
			// need to look to the right of the midPoint
			// -- +1 for the shift to the right
			upperBound = midPoint - 1
		} else if midPointValue == searchVal {
			return true
		}
	}

	return false
}

// try BinarySearch recursion!
func BinarySearchRecur(searchVal int, ints []int) bool {
	// ie, the sliced elements have become too small (therefore no value found!)
	if len(ints) == 0 {
		return false
	}

	lowerBound := 0
	upperBound := len(ints) - 1

	midPoint := (upperBound + lowerBound) / 2
	midPointValue := ints[midPoint]

	if midPointValue < searchVal {
		// need to set lowerbound to right of midPoint
		return BinarySearchRecur(searchVal, ints[midPoint+1:])
	} else if midPointValue > searchVal {
		return BinarySearchRecur(searchVal, ints[:midPoint-1+1]) // added + 1 due to non-inclusion
	} else { // is equal!!
		return true
	}
}

func main() {
	ordered := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(BinarySearch(11, ordered))
	fmt.Println(BinarySearch(88, ordered))
	fmt.Println(BinarySearch(77, ordered))
	fmt.Println(BinarySearch(55, ordered))
	fmt.Println(BinarySearch(44, ordered)) // NlogN calls, true case works
	fmt.Println(BinarySearch(45, ordered)) // false case works

	fmt.Println("----Recursion----")
	fmt.Println(BinarySearchRecur(11, ordered))
	fmt.Println(BinarySearchRecur(88, ordered))
	fmt.Println(BinarySearchRecur(77, ordered))
	fmt.Println(BinarySearchRecur(55, ordered))
	fmt.Println(BinarySearchRecur(44, ordered)) // true case works
	fmt.Println(BinarySearchRecur(45, ordered)) // false case works
}
