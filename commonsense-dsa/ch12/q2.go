package main

import "fmt"

var counter = 0

func GolombMemo(n int, memo map[int]int) int {

	// memo optimization - this is the check
	if val, ok := memo[n]; ok {
		return val
	}

	// memo population @ base case
	if n == 1 {
		memo[1] = 1
		return 1
	}

	counter += 1
	fmt.Println(counter) // <-- for debugging comparison's sake

	// memo population at normal case
	memo[n] = 1 + GolombMemo(n-GolombMemo(GolombMemo(n-1, memo), memo), memo)
	return memo[n]
}

func Golomb(n int) int {
	if n == 1 {
		return 1
	}

	counter += 1
	fmt.Println(counter)

	return 1 + Golomb(n-Golomb(Golomb(n-1)))
}

func AddToMap(myMap map[int]int) {
	myMap[1] = 100
}

func main() {
	// fmt.Println(Golomb(40)) // <--- @40, 1257530 calls

	myMap := make(map[int]int)
	fmt.Println(myMap)
	AddToMap(myMap)
	fmt.Println(myMap) // <-= ok

	fmt.Println(GolombMemo(40, make(map[int]int))) // <--- @40, 39 calls, wow
}
