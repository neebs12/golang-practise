package main

import (
	"fmt"
	"math/rand"
	t "time"
)

func main() {
	rand.Seed(t.Now().Unix()) // <--- for time elapsed in seconds from 1970
	name := "Jason"
	age := 26
	temp := rand.Float32() * 20

	fmt.Printf("I am %v, I am aged %d, and I feel %.2f Celcius \n", name, age, temp)

	statement := fmt.Sprintf("I am %v, I am aged %d, and I feel %.2f Celcius \n", name, age, temp)

	fmt.Println(statement)

	fmt.Println("Hello world")
	fmt.Println(t.Now())
}
