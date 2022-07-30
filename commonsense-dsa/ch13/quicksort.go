package main

import "fmt"

func Partition(leftPtr, rightPtr int, ints []int) int {
	pivotPtr := rightPtr
	pivotVal := ints[pivotPtr]
	rightPtr -= 1 // decrement, as we want rightPtr to the at the imm- left of pivot

	// equiv while loop
	for {
		leftVal := ints[leftPtr]
		rightVal := ints[rightPtr]

		// normal, keep shifting leftPtr to the right
		// -- this clause is ignored if leftVal is now greater than pivot, leftPtr is now static
		// -- second clause to prevent out of index
		for leftVal < pivotVal {
			leftPtr += 1
			leftVal = ints[leftPtr]
		}

		// normal, keep shifting rightPtr to the left
		// -- this clause is ignored if rightVal is lesser than pivot, rightPtr is now static
		// -- second clause to prevent out of index
		for rightVal > pivotVal {
			rightPtr -= 1
			rightVal = ints[rightPtr]
		}

		// what is the break condition here?
		// -- if the pointers have yet to have met, (< used, not <=)
		// -- then we swap their values, to begin next round of movement
		// -- fyi, the left and rightPtrs are guaranteed to be static
		if leftPtr < rightPtr {
			ints[leftPtr], ints[rightPtr] = ints[rightPtr], ints[leftPtr]
			leftPtr += 1
		}

		// break condition is when
		// -- the pointers have met (or gone past), we break
		if leftPtr >= rightPtr {
			break
		}
	}

	// swap the values of the leftPtr and the pivotPtr
	ints[leftPtr], ints[pivotPtr] = ints[pivotPtr], ints[leftPtr]
	return leftPtr
}

func QuickSort(leftPtr, rightPtr int, ints []int) {
	if rightPtr-leftPtr <= 1 {
		return
	}
	// we call the partition on the whole ints
	pivotPtr := Partition(leftPtr, rightPtr, ints)

	QuickSort(leftPtr, pivotPtr-1, ints) // sorting the right hand side

	QuickSort(pivotPtr, rightPtr, ints) // sorting the left hand side
}

func main() {
	foo := []int{0, 9, 1, 2, 6, 3, 4, 5}
	// Partition(0, len(foo)-1, foo)
	QuickSort(0, len(foo)-1, foo)
	fmt.Println(foo)

	// doesnt work with dupliated values :/
	bar := []int{23, 56, 87, 234, 86, 54, 32, 1, 55, 9, 34, 7, 22, 4, 77, 210}
	QuickSort(0, len(bar)-1, bar)
	fmt.Println(bar)
}
