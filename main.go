package main

import (
	"fmt"
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	combined := s1 + s2
	// then these are sorted according to unicode
	sliceCombined := strings.Split(combined, "")
	// then is sorted lexicographically
	// sorted - mutated due to being a slice
	sort.Strings(sliceCombined)
	// then, only concat the string if is not already included
	resultStr := ""
	for _, val := range sliceCombined {
		// val is a single character string
		if !strings.Contains(resultStr, val) {
			resultStr += val
		}
	}
	return resultStr
}

func main() {
	fmt.Println("jason aricheta")
	fmt.Println(TwoToOne("jason", "aricheta"))
}
