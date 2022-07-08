package main

import (
	"fmt"
	"math/rand"
	"time"
)

func greetAndAdd(greet string, n1, n2 int) string {
	return fmt.Sprintf("%s: %d", greet, n1+n2)
}

func myRange(n1, n2 int) (less int, more int) {
	less = n1 - n2
	more = n1 + n2
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i += 1 {
		firstNum := rand.Intn(100)
		secondNum := rand.Intn(100)
		// fmt.Println(greetAndAdd("Hello", firstNum, secondNum))
		fmt.Println(myRange(firstNum, secondNum))
	}
}
