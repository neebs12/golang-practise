package main

import "fmt"

func main() {

	for ind, val := range [3]int{11, 22, 33} {
		fmt.Println(ind, val)
	}

	for ind := range make([]int, 3) {
		fmt.Println(ind)
	}

	mp := map[string]float32{"height": 1.79, "temp": 14.3}
	for key, val := range mp {
		fmt.Println(key, val)
	}

	for _, unicode := range "jason" {
		fmt.Println(unicode)
		fmt.Printf("unicode is of type %T\n", unicode) // int32
	}
}
