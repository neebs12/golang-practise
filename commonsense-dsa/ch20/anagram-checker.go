package main

import "fmt"

func AnagramChecker(s1, s2 string) bool {
	// GCL for unequal lengths, not an annagram
	if len(s1) != len(s2) {
		return false
	}

	myMap := make(map[string]bool)
	for _, rr := range s1 {
		myMap[string(rr)] = false
	}

	// iterate over the s2, flip myMap bool to true for each successful access
	// if there is an unsuccesful access, return false
	// does not account for if s2 has duplicate valid letters, therefore...
	for _, rr := range s2 {
		if _, ok := myMap[string(rr)]; !ok {
			return false
		} else {
			myMap[string(rr)] = true
		}
	}

	// second iteration on myMap, if any are still false, return false
	for _, val := range myMap {
		if !val {
			return false
		}
	}

	// else, at end of the function, return true
	return true
} // is O(M + N) algo

// -- from the book itself
func BetterAnagramChecker(s1, s2 string) bool {
	myMap1 := make(map[byte]int)
	myMap2 := make(map[byte]int)

	for ind, _ := range s1 {
		myMap1[s1[ind]] += 1
		myMap2[s2[ind]] += 1
	}

	for key, val := range myMap1 {
		if val != myMap2[key] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(AnagramChecker("jason", "nosaj"))
	fmt.Println(AnagramChecker("jason", "nojaj"))
	fmt.Println(BetterAnagramChecker("jason", "nosaj"))
	fmt.Println(BetterAnagramChecker("jason", "nojaj"))
}
