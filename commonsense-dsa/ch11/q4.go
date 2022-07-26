package main

import "fmt"

var _ = fmt.Println

func FindIndexForSingleCharStr(large, sub string, ind int) int {
	// condition, `sub` always within `large`
	// in this recur fn,
	// -- large decrementally shortens
	// -- sub stays constant
	// -- ind stores the current index we are in (incr for each call)

	// base case, first element or large[0:1] is == sub
	if charStr := large[0:1]; charStr == sub {
		return ind
	}

	// else sub is not found, increment ind
	ind += 1

	// decrementally 'smallify' large
	large = large[1:]

	// then return
	return FindIndexForSingleCharStr(large, sub, ind)
}

func main() {
	name := "Jason Aricheta"
	fmt.Println(name[1:])  // we can slice strings!
	fmt.Println(name[0])   // we can slice strings!
	fmt.Println(name[0:1]) // we can slice strings!

	fmt.Println("----------------------------")
	fmt.Println(FindIndexForSingleCharStr("jason arichetax", "x", 0))
	fmt.Println(FindIndexForSingleCharStr("xjason aricheta", "x", 0))
	fmt.Println(FindIndexForSingleCharStr("jason arixcheta", "x", 0))
	fmt.Println(FindIndexForSingleCharStr("jason x aricheta", "x", 0))
}
