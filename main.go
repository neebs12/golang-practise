package main

import (
	"fmt"
)

func myCounterFromStart() (func() int, *int) {
	var counter int
	return func() int {
		counter += 1
		return counter
	}, &counter
}

func main() {
	foo := 6
	fmt.Println(&foo, foo)

	var bar *int = &foo
	fmt.Println(bar, foo)
	fmt.Println(&bar, bar, foo)
	*bar = 600
	fmt.Println(foo)

	var baz **int = &bar
	fmt.Println(&bar, bar, foo)
	fmt.Println(baz, bar, foo)

	cnt, cntPtr := myCounterFromStart()

	fmt.Println(cntPtr, cnt())
	fmt.Println(cntPtr, cnt())
	fmt.Println(cntPtr, cnt())

	*cntPtr = 100

	fmt.Println(cntPtr, cnt())
	fmt.Println(cntPtr, cnt())
	fmt.Println(cntPtr, cnt())
}
