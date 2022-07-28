package main

import "fmt"

func FindLongestConsecutive(nums []int) []int {
	// create a map of all numbers
	// also, capture the minimum and maximum of the collection
	intMap := make(map[int]bool)
	min := nums[0]
	max := nums[0]
	for _, val := range nums {
		intMap[val] = true
		if val > max {
			max = val
		} else if val < min {
			min = val
		}
	}

	// init longestSli to make([]int, 0)
	// init currentSli to make([]int, 0)
	longestSli := make([]int, 0)
	currentSli := make([]int, 0)

	// iterate from min to max
	for i := min; i <= max; i += 1 {
		// gcl, if i is not a valid value of intMap, continue
		if _, ok := intMap[i]; !ok {
			continue
		}

		// here, if currentSli length == 0, we append currenSli
		// else if currentSli[len(currentSli) - 1] + 1 == i, append currentSli
		// else, the newly encountered `i` is no longer within sequence
		// -- therefore currentSli = []int{i}, a new slice!
		if len(currentSli) == 0 {
			currentSli = append(currentSli, i)
		} else if currentSli[len(currentSli)-1]+1 == i { // is conseq
			currentSli = append(currentSli, i)
		} else { // conseq is broken, new single elm sli with []int{i}
			currentSli = []int{i}
		}

		// then we compare the length of the currentSli with the longestSli
		// if len(currentSli) > len(longestSli), reassg longestSli
		// -- end iteration
		if len(currentSli) > len(longestSli) {
			longestSli = currentSli
		}
	}
	return longestSli
} // time ON???, but higher space complexity

func main() {
	foo := []int{10, 5, 12, 3, 55, 30, 4, 11, 2}
	bar := []int{19, 13, 15, 12, 18, 14, 17, 11}
	fmt.Println(FindLongestConsecutive(foo))
	fmt.Println(FindLongestConsecutive(bar))
}
