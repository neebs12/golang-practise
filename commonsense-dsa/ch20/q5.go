package main

import "fmt"

// does not sort in place, returns a new slice
func SortBodyTemp(temps []float64) []float64 {
	// conversion from float64 to int to avoid float64 bugs
	tempInts := make([]int, len(temps))
	for ind, val := range temps {
		tempInts[ind] = int(val * 10)
	}

	myMap := make(map[int]int)
	for _, temp := range tempInts {
		myMap[temp] += 1
	}

	minTemp := 970 // are float64
	maxTemp := 990
	sortedTemp := make([]int, 0, len(temps))
	for i := minTemp; i <= maxTemp; i += 1 {
		if nums, ok := myMap[i]; ok {
			counter := 0
			for counter < nums {
				sortedTemp = append(sortedTemp, i)
				counter += 1
			}
		}
	}

	sortedTempFloat64 := make([]float64, len(temps))
	for ind, val := range sortedTemp {
		equivFloat64 := float64(val) / float64(10)
		sortedTempFloat64[ind] = equivFloat64
	}
	return sortedTempFloat64
} // is O(N) + plus annoying float * 10 -> int, then int / 10 -> float64 conversions

func main() {
	temps := []float64{
		98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1,
	}

	fmt.Println(SortBodyTemp(temps))
}
