package main

import "fmt"

func containsX(s string) bool {
	hasX := false
	for _, rr := range s {
		if 'X' == rr {
			hasX = true
		}
	}
	return hasX
}

func main() {
	fmt.Println("Jason", containsX("Jason"))
	fmt.Println("JaXson", containsX("JaXson"))
}
