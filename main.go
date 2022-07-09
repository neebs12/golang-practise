package main

import "fmt"

func whatAmI(val interface{}) {
	switch val := val.(type) {
	case int:
		fmt.Println("I am an integer")
	case bool:
		fmt.Println("I am a boolean")
	case string:
		fmt.Println("I am a string")
	default:
		fmt.Printf("I cant tell what %T is\n", val)
	}
}

func main() {
	whatAmI(100)
	whatAmI(true)
	whatAmI(!!true)
	whatAmI("hello")
	whatAmI(1.1)
}
