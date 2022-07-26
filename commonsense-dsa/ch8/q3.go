package main

import "fmt"

func returnNotIncludedLetter(largeStr string) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	// construct map based on letters of alphabet
	myMap := make(map[rune]bool)
	for _, rr := range alphabet {
		myMap[rr] = true
	}

	// iterate over largeStr
	for _, rr := range largeStr {
		if rr == ' ' {
			continue // ignore spaces
		}
		myMap[rr] = false
	}

	// iterate over myMap, to see which value, for a key has stayed true
	for key, val := range myMap {
		if val == true {
			return string(key)
		}
	}
	return ""
}

func main() {
	fmt.Println(returnNotIncludedLetter("the quick brown box jumps over a lazy do"))
}
