package main

import "fmt"

func main() {
	const foo = 'h'
	const bar = 'e'
	fmt.Printf("Is type %T \n", foo)
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(foo + bar)

	var baz = 4.20
	fmt.Printf("is %T\n", baz)

	fmt.Printf("I am feeling 70%% today\n")
}
