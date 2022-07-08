package main

import (
	"fmt"
)

func main() {
  // try to get name, and add a greeting to it
	var num1, num2 int = 99, 5
	num3 := num1 / num2
	fmt.Printf("the number is %d\n", num3)

	num4 := float32(num1) / float32(num2)
	fmt.Printf("the number is %.2f\n", num4)

	// num5 := num1 / float32(num2) 
	// <--- error due to mismatched types

	var num6 float32 = 3.2
	var num7 int = 3
	// fmt.Println("This is boolean: %t", num6 == num1)
	// <--- error due to mismatched types
	fmt.Printf("This is boolean: %t\n", int(num6) == num7)
}
