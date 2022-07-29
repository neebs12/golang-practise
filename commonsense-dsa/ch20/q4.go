package main

import "fmt"

func GetHighestProduct(ints []int) int {
	// do greedy
	// -- first for min and max
	// -- second, for minMore and maxLess
	// then do prod comparison
	// `greaterProd` = max * maxLess
	// if greaterProd < (min * minMore) make `greaterProd` equal to the min prod

	max := ints[0]
	min := ints[0]
	for _, val := range ints {
		if val > max {
			max = val
		} else if val < min {
			min = val
		}
	}

	maxLess := ints[0]
	minMore := ints[0]
	for _, val := range ints {
		if val > maxLess && val != max {
			maxLess = val
		} else if val < minMore && val != min {
			minMore = val
		}
	}

	minProd := min * minMore
	maxProd := max * maxLess
	greaterProd := maxProd
	if minProd > greaterProd {
		greaterProd = minProd
	}

	return greaterProd
} // O(N) -- two loops, so 2N

func main() {
	foo := []int{5, -10, -6, 9, 4}
	fmt.Println(GetHighestProduct(foo))
}
