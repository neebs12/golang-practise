package main

import (
	"fmt"
	"strconv"
)

func CountBits(num uint) int {
	// convert to binary
	// init `binStr` to ""
	// init `runningInt` to `num`
	// while loop, while runningInt != 0 or 1 (then in loop below)
	// -- divInt = runningInt / 2
	// -- if divInt*2 != runningInt -> binStr += "1"
	// -- else -> binStr += "0"
	// -- finally -> runningInt = divInt
	// -- end iteration
	// once converted to `binStr`, iterate over
	// then Atoi each character str, then sum them together
	// init `summation` to `0`
	// for loop, interate over `binStr` with charBinStr
	// -- equivInt init to strconv.Atoi(charBinStr)
	// -- summation += equivInt
	// -- end iteration
	// return summation

	binStr := ""
	runningInt := num
	for runningInt >= 1 {
		// fmt.Println(runningInt)
		divInt := runningInt / 2
		if divInt*2 != runningInt {
			binStr = "1" + binStr // reversing sequence
		} else {
			binStr = "0" + binStr
		}
		runningInt = divInt // div2 down
	}

	// cannot have a 'final binary num at 0'
	// if binStr[0] == '0' {
	// 	binStr = "1" + binStr
	// }

	summation := 0
	for _, charBinStr := range binStr {
		val, _ := strconv.Atoi(string(charBinStr))
		summation += val
	}

	fmt.Printf("for %d: the final value is %d: the binary is: %s, summation is: %d\n", num, runningInt, binStr, summation)
	// fmt.Println(binStr)

	return summation
}

func main() {
	CountBits(7)
	CountBits(9)
	CountBits(122)
	CountBits(124)
	CountBits(1234)
	CountBits(2)
}
