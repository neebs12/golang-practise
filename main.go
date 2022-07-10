package main

import (
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

// Error method on argError struct required to fulfull the
// -- error interface!
func (a *argError) Error() string {
	return fmt.Sprintf("%d --- %s", a.arg, a.prob)
	// <--- ohh snap, so this is the "message"/"value" assc.
	// -- with the error object!
}

func f2(num int) (int, error) {
	if num == 42 {
		return num, &argError{num, "is not the answer to life!"}
	}
	return num, nil
}

func main() {
	arr := [3]int{11, 42, 89}
	for _, val := range arr {
		if r, e := f2(val); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 success", r)
		}
	}
}
