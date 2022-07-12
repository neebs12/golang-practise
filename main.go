package main

import (
	"math"
)

func GetMiddle(s string) string {
	rr := []rune(s)
	var midIndNum float64 = float64(len(rr)) / 2 // could be n.5(even) or n(odd)
	high, low := math.Ceil(midIndNum), math.Floor(midIndNum)
	if high == low {
		// is odd, can take midIndNum
		return string(s[int(midIndNum)])
	} else {
		return s[int(low) : int(low)+2]
	}
} // trying out: https://www.codewars.com/kata/56747fd5cb988479af000028/train/go

func main() {

}
