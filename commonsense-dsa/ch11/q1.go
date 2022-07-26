package main

import "fmt"

func GetNumberOfCharacters(s []string) int {

	firstVal := s[0]

	// base case, if len of s is == 1
	if len(s) == 1 {
		return len(firstVal)
	}

	return len(firstVal) + GetNumberOfCharacters(s[1:])
}

func main() {
	collection := []string{"ab", "c", "def", "ghij"}
	fmt.Println(GetNumberOfCharacters(collection))
}
