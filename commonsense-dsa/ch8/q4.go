package main

import "fmt"

func returnFirstUnique(str string) string {
	myMap := make(map[rune]int)
	for _, rr := range str {
		myMap[rr] += 1
	}

	// second iteration on str
	for _, rr := range str {
		if val, _ := myMap[rr]; val == 1 {
			return string(rr)
		}
	}
	return ""
}

func main() {
	fmt.Println(returnFirstUnique("minimum"))
}
