package main

import "fmt"

func CountFrom10() {
	for i := 10; i >= 0; i -= 1 {
		fmt.Println(i)
	}
}

func CountdownRecur(number int) {
	fmt.Println(number)
	if number == 0 {
		return
	}

	number -= 1
	CountdownRecur(number)
}

func main() {
	// CountFrom10()
	CountdownRecur(10)
}
