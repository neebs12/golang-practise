package main

import "fmt"

func ReverseCollection(collection []int) {
	for i, j := 0, len(collection)-1; i < j; i, j = i+1, j-1 {
		collection[i], collection[j] = collection[j], collection[i]
	}
}

func main() {
	myCollection := []int{11, 22, 33, 44, 55}
	ReverseCollection(myCollection)
	fmt.Println(myCollection)
}
