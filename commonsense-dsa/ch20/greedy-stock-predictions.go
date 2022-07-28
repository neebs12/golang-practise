package main

import "fmt"

func HasPositiveTrend(nums []float64) bool {
	lowestValue := nums[1]
	middleValue := nums[2]
	for _, val := range nums[3:] {
		// here, val is lower than lowestValue
		if val < lowestValue {
			lowestValue = val

			// here, val is higher than lowestVal but higher than middleValue
		} else if val < middleValue {
			middleValue = val

			// here, we encounter a value that is higher than both lowestVal and middleVal. Therefore completing the trend
		} else {
			return true
		}
	}
	return false
}

func main() {
	foo := []float64{22, 25, 21, 18, 19.8, 17, 16, 20.5}
	fmt.Println(HasPositiveTrend(foo))
	bar := []float64{22, 25, 21, 18, 19.8, 17, 16, 19}
	fmt.Println(HasPositiveTrend(bar))
	baz := []float64{5, 2, 8, 4, 3, 7}
	fmt.Println(HasPositiveTrend(baz))
	qux := []float64{8, 9, 7, 10}
	fmt.Println(HasPositiveTrend(qux))
}
