package main

import "fmt"

func getFirstDuplicate(strCollection []string) string {
	myMap := make(map[string]bool)
	for _, str := range strCollection {
		_, ok := myMap[str]
		if ok {
			return str
		} else {
			myMap[str] = true
		}
	}
	return ""
}

func main() {
	strSli := []string{"a", "b", "c", "d", "c", "e", "f"}
	fmt.Println(getFirstDuplicate(strSli))
}
