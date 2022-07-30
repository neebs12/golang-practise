package main

import "fmt"

// struct
type MySlice struct {
	data []int
}

// Partition function
func (s *MySlice) Partition(leftInd, rightInd int) int {
	pivotInd := rightInd
	pivot := s.data[pivotInd]

	rightInd = pivotInd - 1 // immediate left of the pivot index

	for {
		// keep shifting left ptr to the right if targetted value is greater than `pivot`
		for leftValue := s.data[leftInd]; leftValue < pivot; {
			if leftInd == pivotInd { // leftPtr can point to pivot
				break // out of bounds
			}
			leftInd += 1
			leftValue = s.data[leftInd]
		}

		// keep shifting right ptr to the left if targetted value is less than `pivot`
		for rightValue := s.data[rightInd]; rightValue > pivot; {
			if rightInd == 0 {
				break // out of bounds
			}
			rightInd -= 1
			rightValue = s.data[rightInd]
		}

		// here, we scrutinize the locations of `leftInd` and `rightInd`
		// if the leftInd has moved beyond rightInd, we break the while loop and move on
		if leftInd >= rightInd {
			break
		} else {
			// here the leftInd has encountered a value that is greater than the pivot and the rightInd has encountered a value that is smaller than the pivot, They need to be swapped places
			s.data[leftInd], s.data[rightInd] = s.data[rightInd], s.data[leftInd]
			// then, move the left pivot to the right, for the next round of the immediately wrapping while loop
			leftInd += 1
		}
	}

	// here, we swap the values stored within the left pointer and the pivot
	// so that the pivotInd value is at the correct position within the array
	s.data[leftInd], s.data[pivotInd] = s.data[pivotInd], s.data[leftInd]

	return leftInd
}

// Quicksort function
func (s *MySlice) Quicksort(leftInd, rightInd int) {
	if rightInd-leftInd <= 0 {
		return // left and right ind have approached each other
	}

	pivotInd := s.Partition(leftInd, rightInd)

	s.Quicksort(pivotInd+1, rightInd) // <-- takes care of rhs of pivot
	s.Quicksort(leftInd, pivotInd-1)  // <-- takes care of lhs of pivot
}

func HighestProdFromThreeElms(s MySlice) int {
	s.Quicksort(0, len(s.data)-1)
	return s.data[len(s.data)-1] * s.data[len(s.data)-2] * s.data[len(s.data)-3]
}

func main() {
	mySlice := MySlice{
		data: []int{111, 33, 99, 66, 22, 100, 66, 99, 9},
	}

	fmt.Println(HighestProdFromThreeElms(mySlice))
	fmt.Println(mySlice.data)
}
