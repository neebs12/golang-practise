package main

import "fmt"

func Intersection(coll1, coll2 []int) (intersection []int) {
	myMap := make(map[int]bool)

	for _, val := range coll1 {
		myMap[val] = true
	}

	for _, val := range coll2 {
		_, ok := myMap[val]
		if ok {
			intersection = append(intersection, val)
		}
	}

	return intersection
}

func main() {
	coll1 := []int{1, 2, 3, 4, 5}
	coll2 := []int{0, 2, 4, 6, 8}
	fmt.Println(Intersection(coll1, coll2))
}
