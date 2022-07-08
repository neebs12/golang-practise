package main

import (
	"fmt"
)

func main() {
	// write an algo that (only) prints the duplicates in an array
	sli := []int{11, 4, 44, 564, 654, 11, 10, 6, 4, 4, 9}

	for i, elm := range sli {
		for _, elm2 := range sli[i+1:] {
			if elm == elm2 {
				fmt.Println(elm)
				break
			}
		}
	}
}
