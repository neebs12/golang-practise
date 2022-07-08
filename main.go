package main

import (
	"fmt"
)

func main() {
	var mp map[string]int = map[string]int{
		"jason": 26,
		"val": 25,
		"isaac": 31, 
	}

	fmt.Println(mp, len(mp))
	// <--- notice, unordered
	// <--- notice, use of len is valid

	value, ok := mp["jason"]
	if ok {
		fmt.Printf("%s is %d years old\n", "jason", value)
	}
	// <--- checking if something exists

	mp["jason"] = 27
	fmt.Println(mp["jason"])

	mp["tim"] = 22
	fmt.Println(mp, len(mp))
	// <--- incr in len

	delete(mp, "jason")
	delete(mp, "val")
	delete(mp, "isaac")
	fmt.Println(mp, len(mp))
	// <--- successful deletion

	mp2 := make(map[string]int)
	// <--- another pattern to make maps
	mp2["launch school"] = 10
	mp2["dev academy"] = 8
	mp2["auckland uni"] = 3 
	fmt.Println(mp2)
}
