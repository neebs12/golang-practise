package main

import "fmt"

func Partition(leftPtr, rightPtr int, ints []int) {
	pivotPtr := rightPtr
	pivotVal := ints[pivotPtr]
	rightPtr -= 1 // decrement, as we want rightPtr to the at the imm- left of pivot

	// equiv while loop
	for {
		leftVal := ints[leftPtr]
		rightVal := ints[rightPtr]

		// normal, keep shifting leftPtr to the right
		// -- this clause is ignored if leftVal is now greater than pivot, leftPtr is now static
		if leftVal < pivotVal {
			leftPtr += 1
		}

		// normal, keep shifting rightPtr to the left
		// -- this clause is ignored if rightVal is lesser than pivot, rightPtr is now static
		if rightVal > pivotVal {
			rightPtr -= 1
		}

		// what is the break condition here?
		// -- if the pointers have yet to have met,
		// -- && the rigthPtr is static
		// -- && the leftPtr is static
		// -- then we initiate a swap between the contents of the leftPtr and rightPtr
		if leftPtr < rightPtr && rightVal < pivotVal && leftVal > pivotVal {
			ints[leftPtr], ints[rightPtr] = ints[rightPtr], ints[leftPtr]
		}

		// break condition is when
		// -- the pointers have met (or gone past)
		// -- && the rigthPtr is static
		// -- && the leftPtr is static
		if leftPtr >= rightPtr && rightVal < pivotVal && leftVal > pivotVal {
			// swap the left pointer with the pivot pointer
			ints[leftPtr], ints[pivotPtr] = ints[pivotPtr], ints[leftPtr]
			break
		}
	}
}

func main() {
	foo := []int{0, 5, 2, 1, 6, 3}
	Partition(0, len(foo)-1, foo)
	fmt.Println(foo)
}
